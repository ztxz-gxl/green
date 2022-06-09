<template>
  <div>
    <div class="lyear-layout-web">
      <div class="lyear-layout-container">
        <div>
          <n-link :to="{ name: 'Backindex',params:{id:this,id} }">
            <div class="bottom">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-house-door"
                viewBox="0 0 16 16"
              >
                <path
                  d="M8.354 1.146a.5.5 0 0 0-.708 0l-6 6A.5.5 0 0 0 1.5 7.5v7a.5.5 0 0 0 .5.5h4.5a.5.5 0 0 0 .5-.5v-4h2v4a.5.5 0 0 0 .5.5H14a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.146-.354L13 5.793V2.5a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1.293L8.354 1.146zM2.5 14V7.707l5.5-5.5 5.5 5.5V14H10v-4a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0-.5.5v4H2.5z"
                />
              </svg></div
          ></n-link>
        </div>
        <nav class="navbar navbar-default">
          <div class="topbar">
            <div class="topbar-left">
              <!-- <a href="index.html"><img src="images/logo-sidebar.png" title="LightYear" alt="LightYear" /></a> -->
              <n-link :to="{ name: 'Backindex',params:{id:this,id} }" style="color: black"
                ><span style="font-size: 20px" class="navbar-page-title">
                  后台数据中心
                </span></n-link
              >
            </div>

            <ul class="topbar-right">
              <li class="dropdown dropdown-profile" style="display: flex">
                <a href="javascript:void(0)" data-toggle="dropdown">
                  <img
                    class="img-avatar img-avatar-48 m-r-10"
                    src="images/users/avatar.jpg"
                    alt="绿色生活"
                  />
                  <span style="font-size: 20px"
                    >绿色生活 <span class="caret"></span
                  ></span>
                </a>
                <ul class="dropdown-menu dropdown-menu-right">
                  <li>
                    <n-link :to="{ name: 'Backindex1',params:{id:this,id} }"
                      ><i class="mdi mdi-account"></i> 查看举报信息</n-link
                    >
                  </li>
                  <li>
                    <n-link :to="{ name: 'Backindex4',params:{id:this,id} }"
                      ><i class="mdi mdi-lock-outline"></i>提取银行代币</n-link
                    >
                  </li>
                  <li>
                    <n-link :to="{ name: 'BackLogin',params:{id:this,id} }"
                      ><i class="mdi mdi-lock-outline"></i>退出</n-link
                    >
                  </li>
                </ul>
              </li>
            </ul>
          </div>
        </nav>

        <div class="container-fluid">
          <div class="row">
            <div class="col-sm-6 col-lg-3">
              <div class="card bg-primary">
                <div class="card-body clearfix">
                  <div class="pull-right">
                    <p class="h6 text-white m-t-0">今年中和量</p>
                    <p class="h3 text-white m-b-0">{{ all }}立方米</p>
                  </div>
                  <div class="pull-left">
                    <span class="img-avatar img-avatar-48 bg-translucent"
                      ><svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="currentColor"
                        class="bi bi-droplet"
                        viewBox="0 0 16 16"
                      >
                        <path
                          fill-rule="evenodd"
                          d="M7.21.8C7.69.295 8 0 8 0c.109.363.234.708.371 1.038.812 1.946 2.073 3.35 3.197 4.6C12.878 7.096 14 8.345 14 10a6 6 0 0 1-12 0C2 6.668 5.58 2.517 7.21.8zm.413 1.021A31.25 31.25 0 0 0 5.794 3.99c-.726.95-1.436 2.008-1.96 3.07C3.304 8.133 3 9.138 3 10a5 5 0 0 0 10 0c0-1.201-.796-2.157-2.181-3.7l-.03-.032C9.75 5.11 8.5 3.72 7.623 1.82z"
                        />
                        <path
                          fill-rule="evenodd"
                          d="M4.553 7.776c.82-1.641 1.717-2.753 2.093-3.13l.708.708c-.29.29-1.128 1.311-1.907 2.87l-.894-.448z"
                        /></svg
                    ></span>
                  </div>
                </div>
              </div>
            </div>

            <div class="col-sm-6 col-lg-3">
              <div class="card bg-danger">
                <div class="card-body clearfix">
                  <div class="pull-right">
                    <p class="h6 text-white m-t-0">用户总数</p>
                    <p class="h3 text-white m-b-0">{{ allUser }}人</p>
                  </div>
                  <div class="pull-left">
                    <span class="img-avatar img-avatar-48 bg-translucent"
                      ><svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="currentColor"
                        class="bi bi-person"
                        viewBox="0 0 16 16"
                      >
                        <path
                          d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6zm2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0zm4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4zm-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10c-2.29 0-3.516.68-4.168 1.332-.678.678-.83 1.418-.832 1.664h10z"
                        /></svg
                    ></span>
                  </div>
                </div>
              </div>
            </div>

            <div class="col-sm-6 col-lg-3">
              <div class="card bg-success">
                <div class="card-body clearfix">
                  <div class="pull-right">
                    <p class="h6 text-white m-t-0">参与总次数</p>
                    <p class="h3 text-white m-b-0">{{ join }}次</p>
                  </div>
                  <div class="pull-left">
                    <span class="img-avatar img-avatar-48 bg-translucent"
                      ><svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="currentColor"
                        class="bi bi-check2-circle"
                        viewBox="0 0 16 16"
                      >
                        <path
                          d="M2.5 8a5.5 5.5 0 0 1 8.25-4.764.5.5 0 0 0 .5-.866A6.5 6.5 0 1 0 14.5 8a.5.5 0 0 0-1 0 5.5 5.5 0 1 1-11 0z"
                        />
                        <path
                          d="M15.354 3.354a.5.5 0 0 0-.708-.708L8 9.293 5.354 6.646a.5.5 0 1 0-.708.708l3 3a.5.5 0 0 0 .708 0l7-7z"
                        /></svg
                    ></span>
                  </div>
                </div>
              </div>
            </div>

            <div class="col-sm-6 col-lg-3">
              <div class="card bg-purple">
                <div class="card-body clearfix">
                  <div class="pull-right">
                    <p class="h6 text-white m-t-0">举报总人数</p>
                    <p class="h3 text-white m-b-0">{{ report }} 人</p>
                  </div>
                  <div class="pull-left">
                    <span class="img-avatar img-avatar-48 bg-translucent"
                      ><svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="16"
                        height="16"
                        fill="currentColor"
                        class="bi bi-x-circle"
                        viewBox="0 0 16 16"
                      >
                        <path
                          d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"
                        />
                        <path
                          d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z"
                        /></svg
                    ></span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="row">
            <div class="col-lg-6">
              <div class="card">
                <div class="card-header">
                  <h4>每周用户</h4>
                </div>
                <div class="card-body">
                  <canvas type="2d" class="js-chartjs-bars"></canvas>
                </div>
              </div>
            </div>

            <div class="col-lg-6">
              <div class="card">
                <div class="card-header">
                  <h4>交易历史记录</h4>
                </div>
                <div class="card-body">
                  <canvas type="2d" class="js-chartjs-lines"></canvas>
                </div>
              </div>
            </div>
          </div>

          <div class="row">
            <div class="col-lg-12">
              <div class="card">
                <div class="card-header">
                  <h4>用户数据</h4>
                </div>
                <div class="card-body">
                  <div class="table-responsive">
                    <table style="font-size: 16px" class="table table-hover">
                      <thead>
                        <tr>
                          <th>序号</th>
                          <th>用户地址</th>
                          <th>日期</th>
                          <th>参与方式</th>
                          <th>状态</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="(i, index) in list" :key="index">
                          <td>{{ index + 1 }}</td>
                          <td>{{ i.Addr }}</td>
                          <td>{{ i.Date }}</td>
                          <td>{{ i.Fun }}</td>
                          <td>
                            <span class="label label-success">{{ i.Stu }}</span>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                  <nav
                    aria-label="Page navigation example"
                    style="text-align: center"
                  >
                    <ul class="pagination">
                      <li class="page-item">
                        <n-link :to="{ name: 'Backindex2',params:{id:this,id} }" class="page-link"
                          >查看所有历史记录</n-link
                        >
                      </li>
                    </ul>
                  </nav>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.bottom {
  position: fixed;
  top: 70%;
  left: 80%;
  border: 1px none;
  width: 50px;
  text-align: center;
  height: 50px;
  line-height: 3.8;
  border-radius: 50%;
  background-color: #d4e4d4;
  z-index: 20000000000;
}
</style>

<script>
import axios from "axios";
const { join } = require("path");
export default {
  data() {
    return {
      id: this.$route.params.id,
    };
  },
  head: {
    link: [
      { rel: "stylesheet", href: "css/bootstrap.min.css" },
      { rel: "stylesheet", href: "css/materialdesignicons.min.css" },
      { rel: "stylesheet", href: "css/style.min.css" },
    ],
    script: [
      {
        src: join("/", `js/c.js`),
        defer: true,
      },
    ],
  },
  asyncData() {
    return axios.get(`/api`).then((res) => {
      return {
        all: res.data.all,
        join: res.data.join,
        report: res.data.report,
        allUser: res.data.allUser,
        list: res.data.list,
      };
    });
  },
};
</script>
