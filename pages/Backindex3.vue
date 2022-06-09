<template>
<!-- <div> -->
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
    <div class="row">
            <div class="col-lg-12">
              <div class="card">
                <div class="card-header">
                  <h4>举报信息</h4>
                </div>
                <div class="card-body">
                  <div class="table-responsive">
                    <table style="font-size: 16px" class="table table-hover">
                      <thead>
                        <tr>
                          <th>序号</th>
                          <th>用户地址</th>
                          <th>用户姓名</th>
                          <th>日期</th>
                          <th>是否完成投票</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="(i, index) in list" :key="index">
                          <td>{{ index + 1 }}</td>
                          <td>{{ i.Addr }}</td>
                          <td>{{ i.Name }}</td>
                          <td>{{ i.Date }}</td>
                          <td>{{ i.TFVoting }}</td>
                        </tr>
                      </tbody>
                    </table>
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
const myaxios = axios.create({});
myaxios.interceptors.response.use(
  function (response) {
    return response.data;
  },
  function (error) {
    console.log(error);
  }
);
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
  },
  asyncData() {
    return axios.get(`/api/getRepAll`).then((res) => {
      return {
        list: res.data.list,
      };
    });
  },
};
</script>
