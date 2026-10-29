# Fadewalk 的站点导航

个人名下全部 GitHub Pages 静态站点的导航页。纯静态单文件（`index.html`），无需构建，改完直接推送到 `main` 即可生效。

## 收录的站点

| 分类 | 站点 | 地址 |
|------|------|------|
| 知识库 | AI Infra 可视化学习课程 | https://fadewalk.github.io/ai-infra-visual-course/ |
| 知识库 | AI-Fundamentals 知识库 | https://fadewalk.github.io/AI-fundamentals/ |
| 知识库 | Study · 学习计划（私密） | https://fadewalk.github.io/study/ |
| 项目文档 | KubeOsprey | https://fadewalk.github.io/kube-osprey/ |
| 项目文档 | Ops-CMDB | https://fadewalk.github.io/ops-cmdb/ |
| 项目文档 | 更夫 Gengfu | https://fadewalk.github.io/gengfu/ |
| 图集 | BK-Lite 架构图集 | https://fadewalk.github.io/bk-lite-architecture/ |

## 如何维护

新增站点时编辑 `index.html`，在 `<section id="grid">` 里按现有格式加一张卡片（`data-cat` 取 `learn` / `proj` / `arch`），并同步更新页脚的「更新于」日期。

> 旧的 Foam / MkDocs 笔记模板源码保留在 git 历史与 `bak` 分支中。
