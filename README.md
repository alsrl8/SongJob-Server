### 1. Scheduler server

#### GoLang

It will be running continuously and runs Python crawler once a day, to crawl the web pages.
This means this server has scheduler.
---

### 2. Web Crawler

#### Python, Selenium

It's not running continuously. It runs when it is executed by command-line.
When it runs, it takes arguments of configurations about web crawling.
After finishing web crawling, it will have data from web pages.
It will `produce` to kafka with those data.

---

### 3. Message Queue

#### Kafka

This is message queue handles data from crawler.

---

### 4. Data storing server

#### GoLang

This server is consumer of kafka. It runs continuously and consumes kafka's messages.
If there are new messages, it `consume` them. (Basically the consumed data is from web-crawling)
Now it sends those data to my database cloud.

---

### 5. RESTful API server

#### GoLang, Gin

It handles RESTful API requests. It searches my database cloud for its service.

---

### 6. Database

#### MongoDB, MongoDB Cloud

The database component of this project utilizes MongoDB and MongoDB Cloud, chosen for their flexibility and
user-friendliness in managing data formats and monitoring data. MongoDB, a NoSQL database, is renowned for its ability
to handle large volumes of data and its schema-less nature, which allows for easy modification and evolution of data
structures. This is particularly beneficial in projects where data formats may change frequently. MongoDB Cloud, a
cloud-based database service, offers scalability, high availability, and convenience, ensuring that the database is
accessible and efficient in handling varying workloads.

---

### 7. UI/UX

#### Flutter, Desktop application, Web Application, Mobile Application

The UI/UX component is developed using Flutter, enabling the creation of visually appealing and functionally rich
interfaces across various platforms, including desktop, web, and mobile applications. Flutter's single codebase approach
simplifies development and ensures consistency in design and functionality across different devices. This is
particularly advantageous for maintaining a coherent user experience in multi-platform applications. The use of Flutter
allows for rapid prototyping, quick iterations, and a seamless integration with the backend services, ensuring a
responsive and intuitive user interface that enhances user engagement and satisfaction.
