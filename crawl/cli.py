import argparse

# from crawlers.saramin_crawler import SaraminCrawler
# from crawlers.jobkorea_crawler import JobKoreaCrawler
from crawlers import JumpItCrawler


def parse_arguments():
    parser = argparse.ArgumentParser(description="Run web crawlers for job listings.")
    parser.add_argument("--crawler", choices=["saramin", "jobkorea", "jumpit"], required=True,
                        help="Specify which crawler to run.")
    parser.add_argument("--output", choices=["csv", "json"], default="json",
                        help="Specify the output format.")

    args = parser.parse_args()
    return args


def run_crawler(args):
    crawler = None

    if args.crawler == "saramin":
        pass
        # crawler = SaraminCrawler(options={'output_format': args.output})
    elif args.crawler == "jobkorea":
        pass
        # crawler = JobKoreaCrawler(options={'output_format': args.output})
    elif args.crawler == "jumpit":
        crawler = JumpItCrawler(options={'output_format': args.output})

    if crawler:
        crawler.setup()
        crawler.crawl()
        crawler.teardown()
    else:
        print("Invalid crawler specified.")


def run_cli():
    args = parse_arguments()
    run_crawler(args)
