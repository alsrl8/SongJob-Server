import time
from abc import abstractmethod, ABC

from selenium import webdriver


class BaseCrawler(ABC):
    driver: webdriver.Chrome

    def __init__(self, options=None):
        self.options = options

    @abstractmethod
    def crawl(self):
        pass

    def setup(self):
        self.driver = webdriver.Chrome()

    def teardown(self):
        self.driver.close()

    def scroll_page(self, times, delay):
        for _ in range(times):
            self.driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
            time.sleep(delay)
