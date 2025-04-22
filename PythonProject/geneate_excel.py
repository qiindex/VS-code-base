
import pandas as pd
import json
import os

def excel_to_json(excel_file_path, output_json_path=None):
    """
    读取Excel文件并转换为JSON格式
    
    Args:
        excel_file_path (str): Excel文件的路径
        output_json_path (str, optional): 输出JSON文件的路径。如果为None，则使用Excel文件名
    
    Returns:
        dict: 转换后的JSON数据
    """
    # 读取Excel文件
    df = pd.read_excel(excel_file_path)
    
    # 将DataFrame转换为字典
    data = df.to_dict(orient='records')
    
    # 如果没有指定输出路径，则使用Excel文件名
    if output_json_path is None:
        output_json_path = os.path.splitext(excel_file_path)[0] + '.json'
    
    # 将数据写入JSON文件
    with open(output_json_path, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=4)
    
    return data

if __name__ == "__main__":
    # 示例用法
    excel_file = "/Users/yuanqiang.qi/VsCode/2025-github/new-python-cursor/Account告警准确性治理.xlsx"  # 替换为你的Excel文件路径
    try:
        result = excel_to_json(excel_file)
        print(f"Excel数据已成功转换为JSON格式，并保存到 {os.path.splitext(excel_file)[0] + '.json'}")
    except Exception as e:
        print(f"转换过程中出现错误: {str(e)}")

