# from fabric import *

from fabric import task


@task
def tar_task(c):
    # 打包
    c.run('tar zcvf process-server.tar.gz ./bin/process-server')


@task
def commit(c):
    c.run("git add . && git commit -m '1' ")


@task
def push(c):
    c.run('git push')


@task
def put_task(c):
    # 创建远程服务器文件夹
    with Connection('root@47.107.230.235') as c:
        # 上传文件
        c.run("rm -rf processBack")
        c.run("git clone git@github.com:gouroutong/processBack.git", pty=True)
        c.run('tar zxvf process-server.tar.gz')
        c.put("process-server", "/root/program/process/process-server")
        c.run('supervisorctl update')
        c.run('supervisorctl -c supervisord.conf reload')


@task
def deploy(c):
    tar_task(c)
    commit(c)
    push(c)
    put_task(c)
