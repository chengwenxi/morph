#!/usr/bin/env python3
import os
from github import Github
import anthropic

# 读取 GitHub Token 和 PR 信息
GITHUB_TOKEN = os.getenv("GITHUB_TOKEN")
ANTHROPIC_API_KEY = os.getenv("ANTHROPIC_API_KEY")
PR_NUMBER = os.environ.get("PR_NUMBER")
g = Github(GITHUB_TOKEN)
print(PR_NUMBER)
repo = g.get_repo("chengwenxi/morph")
pr = repo.get_pull(int(PR_NUMBER))  # 获取 PR（你可以根据实际需求获取 PR ID）

# 获取文件列表
files = pr.get_files()

# 提取每个文件的 patch 属性，并排除掉值为 None 的情况
diff = "\n".join([file.patch for file in files if file.patch is not None])

# 生成客户端连接
client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)
# 发送请求
response = client.messages.create(
    model="claude-3-5-sonnet-20240620",
    max_tokens=1024,
    messages=[
        {"role": "user", "content": f"你是一个经验丰富的代码审查助手，负责审查 PR 代码的质量。请审查以下代码变更并提供反馈：\n{diff}"}
    ]
)
# print(response.content)
# 生成评论并提交到 PR
comment = response.content[0].text
pr.create_issue_comment(comment)