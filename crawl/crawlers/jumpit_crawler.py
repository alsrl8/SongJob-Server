import pymongo.errors
from selenium.common.exceptions import NoSuchElementException
from selenium.webdriver.common.by import By

from crawlers import BaseCrawler


class JumpItCrawler(BaseCrawler):
    url = 'https://www.jumpit.co.kr/positions?sort=rsp_rate'

    def __init__(self, options=None):
        super().__init__(options)

    def crawl(self):
        job_infos = []
        self.driver.get(self.url)
        self.scroll_page(0, 3)

        cls_name = 'sc-c8169e0e-0.gOSXGP'
        elements = self.driver.find_elements(By.CLASS_NAME, cls_name)
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

        import repo
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
        job_info['isFavorite'] = False

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
            'company_name': company_name,
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
