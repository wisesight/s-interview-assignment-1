from random import choice
from typing import List
import sys


def check_account_exist(channel_id: str) -> bool:
    return choice((True, False))


def get_new_source() -> List:
    return []


def main():
    new_channels = get_new_source()
    print("=== New Channels ===")
    for channel in new_channels:
        print(
            f"Title: {channel.title} \nsubscriberCount: {channel.subscriber_count}\n\n")
    print("=== End ===")


if __name__ == '__main__':
    sys.exit(main())
