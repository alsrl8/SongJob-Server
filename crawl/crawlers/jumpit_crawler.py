import kafka.errors
import pymongo.errors
from selenium import webdriver
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as ec
from selenium.webdriver.support.ui import WebDriverWait

import repo
from crawlers import BaseCrawler
from message import Producer


class JumpItCrawler(BaseCrawler):
    url = 'https://www.jumpit.co.kr/positions?sort=rsp_rate'
    scroll_num: int

    def __init__(self, scroll_num: int = 0, options=None, headless=True):
        super().__init__(options)
        self.scroll_num = scroll_num
        self.headless = headless

    def setup(self):
        chrome_options = Options()
        if self.headless:
            chrome_options.add_argument('--headless')
        self.driver = webdriver.Chrome(options=chrome_options)

    def crawl(self):
        job_infos = []
        try:
            self.driver.get(self.url)
            body_tag_cls_name = 'sc-c8169e0e-0.gOSXGP'
            WebDriverWait(self.driver, 10).until(ec.presence_of_element_located((By.CLASS_NAME, body_tag_cls_name)))

            self.scroll_page(self.scroll_num, 3)
            elements = self.get_body_tags(body_tag_cls_name)
        except Exception as e:
            print(e)
            return

        for e in elements:
            try:
                job_info = self.extract_job_info(e)
            except NoSuchElementException:
                print(f'Failed to extract job information due to invalid class name')
                continue
            except Exception as e:
                print(e)
                continue

            JumpItCrawler.add_additional_preferences(job_info)
            job_infos.append(job_info)

        self.extract_job_info_detail(job_infos)
        JumpItCrawler.send_to_message_queue(job_infos)

    @staticmethod
    def send_to_message_queue(job_infos):
        print(f'Sending crawled job infos to kafka container. {len(job_infos)=}')
        try:
            producer = Producer()
        except kafka.errors.NoBrokersAvailable:
            print("Kafka container isn't running")
            return

        for job_info in job_infos:
            producer.send(job_info)
        producer.close()
        print(f'Send job posts to message queue. {len(job_infos)=}')

    def get_body_tags(self, cls_name):
        try:
            elements = self.driver.find_elements(By.CLASS_NAME, cls_name)
        except NoSuchElementException as e:
            print(f"Can't find element with class name({cls_name})")
            raise e
        except Exception as e:
            return e
        return elements

    @staticmethod
    def store_crawled_data_to_mongo(job_infos):
        db_name = 'job'
        coll_name = 'job_info'
        client = repo.mongo.get_mongo_client()
        for job_info in job_infos:
            try:
                repo.mongo.insert_document(client, db_name, coll_name, job_info)
            except pymongo.errors.DuplicateKeyError:
                print(f"There is already duplicated job info with same link({job_info['link']})")
                continue
        repo.mongo.close_mongo_client(client)

    @staticmethod
    def add_additional_preferences(job_info):
        job_info['hasBeenViewed'] = False

    def extract_job_info_detail(self, job_infos):
        for job_info in job_infos:
            link = job_info['link']
            self.driver.get(link)
            parent_element = self.driver.find_element(By.TAG_NAME, 'Body')

            try:
                JumpItCrawler.extract_job_tags(job_info, parent_element)
            except Exception as e:
                print(e)
                continue

            try:
                JumpItCrawler.extract_techniques(job_info, parent_element)
            except Exception as e:
                print(e)
                continue

    @staticmethod
    def extract_techniques(job_info, parent_element):
        cls_name = 'sc-3bc54f04-0.bIaunD'
        try:
            tech_infos = parent_element.find_elements(By.CLASS_NAME, cls_name)
        except NoSuchElementException as e:
            print(f"Can't find element with class name({cls_name})")
            raise e
        except Exception as e:
            raise e
        techniques = [ti.text.strip() for ti in tech_infos]
        job_info['techniques'] = techniques

    @staticmethod
    def extract_job_tags(job_info, parent_element):
        cls_name = 'position_tags'
        try:
            tags = JumpItCrawler.extract_text_list(parent_element, cls_name)
        except NoSuchElementException:
            job_info['tags'] = []
            return
        except Exception as e:
            raise e
        job_info['tags'] = tags

    @staticmethod
    def extract_job_info(element):
        tag_name = 'a'
        try:
            link = JumpItCrawler.extract_link(element, tag_name)
        except Exception as e:
            raise e

        cls_name = 'sc-2525ae7a-0.fIwSBa'
        try:
            info_tag = element.find_element(By.CLASS_NAME, cls_name)
        except NoSuchElementException as e:
            print(f"Can't find element with class name({cls_name})")
            raise e
        except Exception as e:
            raise e

        try:
            company_name_cls_name = 'sc-2525ae7a-2.leumnv'
            company_name = JumpItCrawler.extract_text(info_tag, company_name_cls_name)
        except Exception as e:
            raise e

        try:
            name_cls_name = 'position_card_info_title'
            name = JumpItCrawler.extract_text(info_tag, name_cls_name)
        except Exception as e:
            raise e

        try:
            other_info_cls_name = 'sc-2525ae7a-1.hbojxM'
            other_info_list = JumpItCrawler.extract_text_list(info_tag, other_info_cls_name)
        except Exception as e:
            raise e

        if len(other_info_list) != 2:
            raise ValueError("The length of other_info_list must be exactly 2.")

        location = other_info_list[0]
        career = other_info_list[1]

        return {
            'link': link,
            'name': name,
            'company': company_name,
            'location': location,
            'career': career
        }

    @staticmethod
    def extract_link(element, tag_name):
        try:
            link_tag = element.find_element(By.TAG_NAME, tag_name)
            link = link_tag.get_attribute('href')
            return link
        except NoSuchElementException as e:
            print(f"Can't find element with tag name({tag_name})")
            raise e

    @staticmethod
    def extract_text(info_tag, cls_name):
        try:
            text = info_tag.find_element(By.CLASS_NAME, cls_name).text.strip()
        except NoSuchElementException as e:
            print(f"Can't find element with class name({cls_name})")
            raise e
        except Exception as e:
            print(f'Exception occurred extracting text from jumpit page')
            raise e
        return text

    @staticmethod
    def extract_text_list(info_tag, cls_name):
        try:
            element = info_tag.find_element(By.CLASS_NAME, cls_name)
        except NoSuchElementException as e:
            print(f"Can't find element with class name({cls_name})")
            raise e
        except Exception as e:
            print(f'Exception occurred extracting text from jumpit page')
            raise e

        tag_name = 'li'
        try:
            li_elements = element.find_elements(By.TAG_NAME, tag_name)
            return [li.text for li in li_elements]
        except NoSuchElementException as e:
            print(f"Can't find element with tag name({tag_name})")
            raise e
        except Exception as e:
            print(f'Exception occurred extracting text from jumpit page')
            raise e
