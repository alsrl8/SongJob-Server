import argparse

# from crawlers.saramin_crawler import SaraminCrawler
# from crawlers.jobkorea_crawler import JobKoreaCrawler
from crawlers import JumpItCrawler


def parse_arguments():
    parser = argparse.ArgumentParser(description="Run web crawlers for job listings.")
    parser.add_argument("--scroll-num", type=int, default=0,
                        help="Specify the number of times to scroll (integer).")
    parser.add_argument("--crawler", choices=["saramin", "jobkorea", "jumpit"], required=True,
                        help="Specify which crawler to run.")
    parser.add_argument("--output", choices=["csv", "json"], default="json",
                        help="Specify the output format.")
    parser.add_argument("--headless", type=bool, default=False,
                        help="Set to True to run the browser in headless mode (without a GUI), or False to run with a GUI.")

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
        crawler = JumpItCrawler(scroll_num=args.scroll_num, options={'output_format': args.output}, headless=args.headless)

    if crawler:
        crawler.setup()
        crawler.crawl()
        crawler.teardown()
    else:
        print("Invalid crawler specified.")


def run_cli():
    args = parse_arguments()
    run_crawler(args)
