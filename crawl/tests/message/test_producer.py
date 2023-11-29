import unittest
from message import Producer


class MongoTest(unittest.TestCase):

    def test_send(self):
        producer = Producer('mock-topic')
        data = {
            'header': 'Send Test',
            'content': 'This is test',
        }
        producer.send(data)
        producer.close()
