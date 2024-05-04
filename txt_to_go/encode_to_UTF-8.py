import os
import chardet
import codecs

i = 0

def convert_encoding(file_path, target_encoding='utf-8'):
    global i
    i += 1
    # Detect current encoding
    with open(file_path, 'rb') as f:
        rawdata = f.read()
    current_encoding = chardet.detect(rawdata)['encoding']
    print(i,": befor_current_encoding :",current_encoding)

    # Open the file in the detected encoding
    with codecs.open(file_path, 'r', encoding=current_encoding, errors='ignore') as f:
        content = f.read()

    # Write the content back to the file in the target encoding
    with codecs.open(file_path, 'w', encoding=target_encoding) as f:
        f.write(content)
        
    # Chack current encoding
    with open(file_path, 'rb') as f:
        rawdata = f.read()
    current_encoding = chardet.detect(rawdata)['encoding']
    print(i, ": after_current_encoding :",current_encoding)

def convert_encodings_in_folder(folder_path):
    for root, dirs, files in os.walk(folder_path):
        for file in files:
            if file.endswith('.txt'):
                file_path = os.path.join(root, file)
                convert_encoding(file_path)

# 基準フォルダパスを指定してください
folder_path = '.'
convert_encodings_in_folder(folder_path)
