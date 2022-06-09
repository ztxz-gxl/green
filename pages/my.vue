<template>
  <div>
    <div class="navbar-nav w-100"></div>

    <div class="content">
      <nav
        class="navbar navbar-expand bg-light navbar-light sticky-top px-4 py-0"
      >
        <n-link class="btn btn-sm btn-primary mynl" :to="{ path: '/'}"
          >主页</n-link
        >
        <n-link :to="{ name: 'index7' }" class="btn btn-sm btn-primary mynl"
          >退出</n-link
        >
        <div class="navbar-nav align-items-center ms-auto">
          <div class="nav-item dropdown">
            <span class="d-none d-lg-inline-flex">HELLO</span>
            <span class="d-none d-lg-inline-flex">{{user.PersonalName}}</span>
          </div>
        </div>
      </nav>

      <div class="container-fluid pt-4 px-4">
        <div class="row g-4">
          <div class="col-sm-6 col-xl-3">
            <div
              class="
                bg-light
                rounded
                d-flex
                align-items-center
                justify-content-between
                p-4
              "
            >
              <i class="fa fa-user fa-3x text-primary"></i>
              <div class="ms-3" style="text-align: end;">
                <p class="mb-2">用户名</p>
                <h2 class="mb-0" style="font-size: 0.9rem;">{{user.PersonalName}}</h2>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-xl-3">
            <div
              class="
                bg-light
                rounded
                d-flex
                align-items-center
                justify-content-between
                p-4
              "
            >
              <i class="fa fa-chart-line fa-3x text-primary"></i>
              <div class="ms-3" style="text-align: end;">
                <p class="mb-2">被举报次数</p>
                <h2 class="mb-0" style="font-size: 0.9rem;">{{user.NumberOfReported}}</h2>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-xl-3">
            <div
              class="
                bg-light
                rounded
                d-flex
                align-items-center
                justify-content-between
                p-4
              "
            >
              <i class="fa fa-chart-area fa-3x text-primary"></i>
              <div class="ms-3" style="text-align: end;">
                <p class="mb-2">累积获得总能量</p>
                <h2 class="mb-0" style="font-size: 0.9rem;">{{user.NeutralizedAmount}}</h2>
              </div>
            </div>
          </div>
          <div class="col-sm-6 col-xl-3">
            <div
              class="
                bg-light
                rounded
                d-flex
                align-items-center
                justify-content-between
                p-4
              "
            >
              <i class="fa fa-chart-pie fa-3x text-primary"></i>
              <div class="ms-3 " style="text-align: end;overflow: auto;word-break: keep-all;">
                <p class="mb-2">代币数量</p>
                <h2 class="mb-0" style="font-size: 0.9rem;">{{user.PersonalAmount}}</h2>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="container-fluid pt-4 px-4">
        <div class="row g-4">
          <div class="col-sm-12 col-xl-6">
            <div class="bg-light text-center rounded p-4">
              <div
                class="d-flex align-items-center justify-content-between mb-4"
              >
                <h6 class="mb-0">代币奖励</h6>
              </div>
              <canvas id="worldwide-sales"></canvas>
            </div>
          </div>
          <div class="col-sm-12 col-xl-6">
            <div class="bg-light text-center rounded p-4">
              <div
                class="d-flex align-items-center justify-content-between mb-4"
              >
                <h6 class="mb-0">中和能量</h6>
              </div>
              <canvas id="salse-revenue"></canvas>
            </div>
          </div>
        </div>
      </div>

      <div class="container-fluid pt-4 px-4">
        <div class="bg-light text-center rounded p-4">
          <div class="d-flex align-items-center justify-content-between mb-4">
            <h6 class="mb-0">最近获得的能量</h6>
            <n-link :to="{path:'/myALL'}">展示所有</n-link>
          </div>
          <div class="table-responsive">
            <table
              class="
                table
                text-start
                align-middle
                table-bordered table-hover
                mb-0
              "
              style="word-break: keep-all; text-align: center !important"
            >
              <thead>
                <tr class="text-dark">
                  <th scope="col">时间</th>
                  <th scope="col">行为</th>
                  <th scope="col">状态</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(i, index) in list" :key="index">
                  <td>{{i.Time}}</td>
                  <td>{{i.Fun}}</td>
                  <td>EN:{{i.Stu}}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
    <!-- <a href="#" class="btn btn-lg btn-primary btn-lg-square back-to-top"><svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        fill="currentColor"
        class="bi bi-arrow-up"
        viewBox="0 0 16 16"
      >
        <path
          fill-rule="evenodd"
          d="M8 15a.5.5 0 0 0 .5-.5V2.707l3.146 3.147a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L7.5 2.707V14.5a.5.5 0 0 0 .5.5z"
        /></svg
    ></a> -->
  </div>
</template>

<style>
.mynl {
  margin-left: 2%;
  margin-right: 2%;
}
</style>

<script>
import axios from "axios";
const { join } = require("path");
export default {
  data(){
    return{
    user:[],
    list:[],
    }  
    },
  mounted() {
    axios.get(`/api/getUser`).then((res) => {
      Array.from(res.data.list).forEach(i => {
        i.Time = new Date(parseInt(res.data.list[0].Time) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ')
        if(i.Stu){
          i.Stu = "完成"
        }else{
          i.Stu = "被举报"
        }
      }
      )
        this.user=res.data.user
        this.list=res.data.list
    });
  },
  head: {
    link: [
      { rel: "preconnect", href: "https://fonts.googleapis.com/" },
      { rel: "preconnect", href: "https://fonts.gstatic.com/" },
      { rel: "stylesheet", href: "css/acs/css/bootstrap.min.css" },
      { rel: "stylesheet", href: "css/acs/css/all.min.css" },
      { rel: "stylesheet", href: "css/acs/css/style.css" },
      { rel: "stylesheet", href: "css/acs/css/owl.carousel.min.css" },
      {
        rel: "stylesheet",
        href: "css/acs/css/tempusdominus-bootstrap-4.min.css",
      },
      { rel: "stylesheet", href: "css/acs/css/bootstrap-icons.css" },
    ],
    script: [
      { src: join("/", `js/ajs/jquery-3.6.0.min.js`), defer: true },
      { src: join("/", `js/ajs/bootstrap.bundle.min.js`), defer: true },
      { src: join("/", `js/ajs/js/easing.min.js`), defer: true },
      { src: join("/", `js/ajs/js/waypoints.min.js`), defer: true },
      { src: join("/", `js/ajs/js/owl.carousel.min.js`), defer: true },
      { src: join("/", `js/ajs/js/moment.min.js`), defer: true },
      { src: join("/", `js/ajs/js/main.js`), defer: true },
      { src: join("/", `js/ajs/js/rocket-loader.min.js`), defer: true },
    ],
  },
};
</script>

