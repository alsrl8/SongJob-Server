import json
from datetime import datetime

import kafka

from message import get_kafka_url


class Producer:
    server_url: str
    topic: str
    _produce: kafka.KafkaProducer

    def __init__(self, topic: str = 'job_posts'):
        self.server_url = get_kafka_url()
        self.topic = topic + '_' + datetime.now().strftime('%y%m%d')
        self._producer = kafka.KafkaProducer(
            bootstrap_servers=self.server_url,
            value_serializer=lambda v: json.dumps(v).encode('utf-8'),
        )

    def send(self, data: dict):
        try:
            self._producer.send(
                self.topic,
                value=data
            )
        except Exception as e:
            print(f'Error sending message: {e}')
        print(f'Send message to message queue')
        print(f'{data=}')

    def close(self):
        self._producer.close()
        print(f'Close kafka producer')
