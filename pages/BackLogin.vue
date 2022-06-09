<template>
  <div class="container-fluid position-relative bg-white d-flex p-0">
    <div class="container-fluid">
      <div
        class="row h-100 align-items-center justify-content-center"
        style="min-height: 100vh"
      >
        <div class="col-12 col-sm-8 col-md-6 col-lg-5 col-xl-4">
          <div class="bg-light rounded p-4 p-sm-5 my-4 mx-3">
            <div class="d-flex align-items-center justify-content-between mb-3">
              <h1>登录数据中心</h1>
            </div>
            <div>
              <p>用户名会关联您的账户地址</p>
            </div>
            <hr />
            <div class="form-floating mb-3">
              <input
                type="text"
                class="form-control"
                name="name"
                id="floatingText1"
                placeholder="jhondoe"
              />
              <label for="floatingText">请输入用户名</label>
            </div>
            <div class="form-floating mb-3">
              <input
                type="text"
                class="form-control"
                name="password"
                id="floatingText2"
                placeholder="jhondoe"
              />
              <label for="floatingText">请输入账户密码</label>
            </div>
            <div>
              <button type="button" class="btn btn-primary rounded-pill m-2">
                取消
              </button>
              <button
                type="button"
                @click="test"
                class="btn btn-primary rounded-pill m-2"
              >
                登录
              </button>
              <router-link
                :to="{path:'/upload'}"
                type="button"
                class="btn btn-primary rounded-pill m-2"
              >
                绑定
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  head: {
    link: [
      { rel: "stylesheet", href: "css/acs/css2.css" },
      { rel: "stylesheet", href: "static/css/bootstrap-icons.css" },
      { rel: "stylesheet", href: "css/acs/owl.carousel.min.css" },
      {
        rel: "stylesheet",
        href: "static/css/tempusdominus-bootstrap-4.min.css",
      },
      { rel: "stylesheet", href: "css/acs/bootstrap.min.css" },
      { rel: "stylesheet", href: "static/css/style.css" },
    ],
    script: [
      {
        src: "static/js/jquery-3.4.1.min.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      {
        src: "static/js/easing.min.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      {
        src: "static/js/waypoints.min.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      {
        src: "static/js/moment.min.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      {
        src: "static/js/moment-timezone.min.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      {
        src: "static/js/tempusdominus-bootstrap-4.min.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      {
        src: "static/js/main.js",
        type: "51e0fbf48a5089fe5b05817f-text/javascript",
        defer: true,
      },
      { src: "static/js/rocket-loader.min.js", defer: true },
    ],
  },
  methods: {
    test() {
      let username = document.getElementById("floatingText1").value;
      let password = document.getElementById("floatingText2").value;
      return axios
        .get(`/api/login?username=` + username + "&password=" + password)
        .then((res) => {
          if (res.status == 200) {
            this.$router.push({
              name: `Backindex`,
              params: { id: res.data.id },
            });
          }
        });
    },
  },
};
</script>