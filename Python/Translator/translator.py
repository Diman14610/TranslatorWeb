import argparse
import base64
from googletrans import Translator

if __name__ != '__main__':
    exit()

parser = argparse.ArgumentParser()
parser.add_argument('-t', help='text')
parser.add_argument('-s', help='source text language', default='en')
parser.add_argument('-d', help='translating to language', default='ru')

args = parser.parse_args()

if not args.t:
    print('-t empty')
    exit()

translator = Translator()
result = translator.translate(text=args.t, src=args.s, dest=args.d)

print(base64.b64encode(result.text.encode()).decode())
