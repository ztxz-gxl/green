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
              <h1>导入账户</h1>
            </div>
            <div>
              <p>导入的账户将与geth账户相关联。</p>
            </div>
            <hr />
            <form action="/api/upload" method="post" enctype="multipart/form-data">
              <div class="mb-3">
                <input
                  class="form-control"
                  name="upload"
                  type="file"
                  id="formFileMultiple"
                  multiple=""
                />
              </div>
              <div class="form-floating mb-3">
                <input
                  type="text"
                  class="form-control"
                  name="name"
                  id="floatingText"
                  placeholder="jhondoe"
                />
                <label for="floatingText">请设置用户名</label>
              </div>
              <div class="form-floating mb-3">
                <input
                  type="password"
                  class="form-control"
                  name="pas"
                  id="floatingText"
                  placeholder="jhondoe"
                />
                <label for="floatingText">请输入geth密码</label>
              </div>
              <div class="form-floating mb-3">
                <input
                  type="password"
                  class="form-control"
                  name="password"
                  id="floatingText"
                  placeholder="jhondoe"
                />
                <label for="floatingText">请输入账户密码</label>
              </div>
              <div>
                <router-link :to="{path:'/index7'}" type="button" class="btn btn-primary rounded-pill m-2">
                  取消
                </router-link>
                <button type="submit" class="btn btn-primary rounded-pill m-2">
                  导入
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
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
};
</script>
