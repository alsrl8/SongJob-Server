import unittest
import repo


class MongoTest(unittest.TestCase):
    def test_get_mongo_client(self):
        client = repo.mongo.get_mongo_client()
        self.assertIsNotNone(client)
        repo.mongo.close_mongo_client(client)

    def test_insert_document(self):
        client = repo.mongo.get_mongo_client()
        self.assertIsNotNone(client)
        repo.mongo.insert_document(client, 'dummy', 'dummy_coll', {'dummy_key': 'This is dummy.'})
        repo.mongo.close_mongo_client(client)
