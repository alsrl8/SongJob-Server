import os
from typing import List, Dict

from pymongo import MongoClient


def get_cloud_url():
    return os.environ.get('X-SongUser-MongoCloud-Url')


def get_mongo_client():
    url = get_cloud_url()
    client = MongoClient(url)
    return client


def close_mongo_client(client: MongoClient):
    client.close()


def insert_document(client: MongoClient, db_name: str, coll_name: str, doc: dict):
    client[db_name][coll_name].insert_one(doc)


def insert_documents(client: MongoClient, db_name: str, coll_name: str, docs: List[Dict]):
    client[db_name][coll_name].insert_many(docs)
