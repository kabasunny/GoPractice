import os
import chardet

def detect_encoding(file_path):
    with open(file_path, 'rb') as f:
        rawdata = f.read()
    return chardet.detect(rawdata)['encoding']

def detect_encodings_in_folder(folder_path):
    for root, dirs, files in os.walk(folder_path):
        for file in files:
            if file.endswith('.txt'):
                file_path = os.path.join(root, file)
                encoding = detect_encoding(file_path)
                print(f'File: {file_path}, Encoding: {encoding}')

# 基準フォルダパスを指定してください
folder_path = '.'
detect_encodings_in_folder(folder_path)
