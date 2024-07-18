import os

# specify the directory path
# dir_path = 'K:\Коммерческий отдел'

# get a list of all files in the directory
# files = [f for f in os.listdir(dir_path) if os.path.isfile(os.path.join(dir_path, f))]

import os

res = {'text' : 0, 'image':0, 'document' : 0, 'executable' : 0, 'binary': 0}

def get_files_and_types(root_dir):
    for root, dirs, files in os.walk(root_dir):
        for file in files:
            file_path = os.path.join(root, file)
            file_type = None
            
            # check if the file is a text file
            if file.endswith(('.txt', '.csv', '.log')):
                res['text'] += 1
                # file_type = 'text'
            
            # check if the file is an image file
            elif file.endswith(('.jpg', '.jpeg', '.png', '.gif', '.bmp')):
                # file_type = 'image'
                res['image'] += 1
            
            # check if the file is a document file
            elif file.endswith(('.doc', '.docx', '.pdf', '.xls', '.xlsx')):
                # file_type = 'document'
                res['document'] += 1
            
            # check if the file is an executable file
            elif file.endswith(('.exe', '.bin', '.sh')):
                # file_type = 'executable'
                res['executable'] += 1
            
            # if none of the above, assume it's a binary file
            else:
                # file_type = 'binary'
                res['binary'] += 1
            
            # print(f"File: {file_path}, Type: {file_type}")
    for k,v in res:
        print(f"{k}: {v}")

# specify the root directory
root_dir = 'K:\Коммерческий отдел'

get_files_and_types(root_dir)

# def main():
#     pass

# if __name__ == '__main__':
#     main()