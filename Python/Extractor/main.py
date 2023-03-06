import argparse
import base64
from pathlib import Path
import pytesseract
from PIL import Image

if __name__ != '__main__':
    exit()

parser = argparse.ArgumentParser()
parser.add_argument('-f', help='path to file')
parser.add_argument('-o', help='output')
parser.add_argument('-e', help='encoding', default='utf-16')
parser.add_argument('-l', help='language', default='rus')

args = parser.parse_args()

if not args.f:
    print('-f empty')
    exit()
if not Path(args.f).exists():
    print('not exists')
    exit()

img = Image.open(args.f)

pytesseract.pytesseract.tesseract_cmd = 'lib/Tesseract/tesseract.exe'
text: str = pytesseract.image_to_string(
    img,
    lang=args.l,
    config='--oem 3 --psm 6'
)

print(base64.b64encode(text.encode()).decode())

if not args.o:
    exit()
if not Path(args.o).is_dir():
    print('not folder')
    exit()

with open(f'{args.o}/{Path(args.f).stem}.txt', 'w', encoding=args.e) as text_file:
    text_file.write(text)
