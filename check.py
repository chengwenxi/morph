#!/usr/bin/env python3
import os
import openai
from github import Github
# 读取 GitHub Token 和 PR 信息
GITHUB_TOKEN = os.getenv("GITHUB_TOKEN")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
PR_NUMBER = os.environ.get("PR_NUMBER")
g = Github(GITHUB_TOKEN)
print(PR_NUMBER)
repo = g.get_repo("chengwenxi/morph")
pr = repo.get_pull(int(PR_NUMBER))  # 获取 PR（你可以根据实际需求获取 PR ID）
# 获取 PR 代码变更
diff = "\n".join([file.patch for file in pr.get_files()])
# 调用 OpenAI API 进行代码审查
openai.api_key = OPENAI_API_KEY
response = openai.ChatCompletion.create(
    model="anthropic/claude-3-5-sonnet-20240620",
    messages=[
        {"role": "system", "content": "你是一个经验丰富的代码审查助手，负责审查 PR 代码的质量。"},
        {"role": "user", "content": f"请审查以下代码变更并提供反馈：\n{diff}"}
    ]
)
# 生成评论并提交到 PR
comment = response["choices"][0]["message"]["content"]
pr.create_issue_comment(comment)