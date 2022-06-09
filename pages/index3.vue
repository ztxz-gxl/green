<template>
  <div>
    <!-- <div class="preloader">
      <img class="preloader__image" width="60" src="" alt="" />
    </div> -->
    <!-- /.preloader -->
    <div class="page-wrapper">
      <header class="main-header clearfix">
        <nav class="main-menu clearfix">
          <div class="main-menu-wrapper clearfix">
            <div class="main-menu-wrapper-inner clearfix">
              <div class="main-menu-wrapper__left clearfix">
                <div class="main-menu-wrapper__logo">
                  <n-link :to="{ path: '/' }"
                    ><img src="" alt=""
                  /></n-link>
                </div>
                <div class="main-menu-wrapper__main-menu">
                  <a href="#" class="mobile-nav__toggler"
                    ><i class="fa fa-bars"></i
                  ></a>
                 <ul class="main-menu__list">
                    <li>
                      <n-link :to="{ path: '/' }"
                        >首页</n-link
                      >
                    </li>
                    <li>
                      <n-link :to="{ path: '/index2' }"
                        >服务</n-link
                      >
                    </li>
                    <li>
                      <n-link :to="{ path: '/index3' }"
                        >数据中心</n-link
                      >
                    </li>

                    <li>
                      <n-link :to="{ path: '/index4' }"
                        >绿色银行</n-link
                      >
                    </li>
                    <li>
                      <n-link :to="{ path: '/report' }"
                        >举报中心</n-link
                      >
                    </li>
                    <li class="dropdown">
                        <a href="#">钱包</a>
                        <ul>
                            <li><n-link :to="{path:'/approve'}">代币授权</n-link></li>
                            <li><n-link :to="{path:'/transfer'}">代币转账</n-link></li>
                            <li><n-link :to="{path:'/info'}">代币信息</n-link></li>
                      
                        </ul>
                    </li>
                    <li>
                      <n-link :to="{ path: '/my'}"
                        >个人中心</n-link
                      >
                    </li>
                    <li>
                      <n-link :to="{ path: '/index7'}"
                        >切换账户</n-link
                      >
                    </li>
                    <li>
                      <n-link :to="{ path: '/index7'}"
                        >退出</n-link
                      >
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </nav>
      </header>

      <div class="stricky-header stricked-menu main-menu">
        <div class="sticky-header__content"></div>
      </div>
      <section class="about-two" style="padding: 0">
        <div class="container">
          <div class="row" style="height: 465px">
            <div class="col-xl-6">
              <div class="col-sm-6" style="max-width: 90%">
                <h2 class="text-center">用户绿色行为排行榜</h2>
                <canvas id="pie-graph"></canvas>
              </div>
            </div>
            <div class="col-xl-6">
              <div class="about-two__right" style="margin-left: auto">
                <div class="section-title text-left">
                  <span class="section-title__tagline">用户绿色行为排行榜</span>
                </div>
                <p class="about-two__text-1">
                  您可以在这查看今日参加的绿色生活行为最多的用户，这样您可以拥有更直观的感受，恭喜这些用户为环境做出的贡献，快为我们的美好环境一起努力!!希望下一次的榜一会是你
                </p>
              </div>
            </div>
          </div>
          <div class="row" style="margin-top: 10%; margin-bottom: 10%">
            <div class="col-xl-6">
              <div class="col-sm-6" style="max-width: 100%">
                <h2 class="text-center">绿色行为方式排行榜</h2>
                <canvas id="doughnut-graph"></canvas>
              </div>
            </div>
            <div class="col-xl-6">
              <div class="about-two__right" style="margin-left: auto;margin-top: 50px;">
                <div class="section-title text-left">
                  <span class="section-title__tagline">绿色行为方式排行榜</span>
                </div>
                <p class="about-two__text-1">
                  您可以在这查看今日最火的绿色生活方式，这里也有你做的一份贡献，这样您可以拥有更直观的感受，快为我们的美好环境一起努力!!希望下一次的榜一会是你
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>
       <section class="contact-page">
        <div class="container">
          <div class="row" style="border: solid; border-width: 1px;border-radius: 5px">
            <div id="cnm">
              <p>被举报用户</p>
            </div>
            <table style="text-align: center;word-break: keep-all;">
              <thead style="background: aquamarine; color: black">
                <tr>
                  <th>序号</th>
                  <th>用户名</th>
                  <th>截止时间</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody style="background-color: aliceblue">
                <tr v-for="(i, index) in list" :key="index">
                  <td>{{ index + 1 }}</td>
                  <td>{{ i.Name }}</td>
                  <td>{{ i.VotingTime }}</td>
                  <td>
                    <a class="btn btn-sm btn-primary" href="javascript:void(0)" @click="vote(index + 1)"
                      >投票</a
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </section>
    </div>
    <!-- /.page-wrapper -->

    <div class="mobile-nav__wrapper">
      <div class="mobile-nav__overlay mobile-nav__toggler"></div>
      <!-- /.mobile-nav__overlay -->
      <div class="mobile-nav__content">
        <span class="mobile-nav__close mobile-nav__toggler"
          ><i class="fa fa-times"></i
        ></span>

        <div class="logo-box">
          <n-child :to="{ path: '/' }" aria-label="logo image"
            ><img src="" width="155" alt=""
          /></n-child>
        </div>
        <!-- /.logo-box -->
        <div class="mobile-nav__container"></div>
        <!-- /.mobile-nav__container -->
      </div>
      <!-- /.mobile-nav__content -->
    </div>
    <!-- /.mobile-nav__wrapper -->

    <div class="search-popup">
      <div class="search-popup__overlay search-toggler"></div>
      <!-- /.search-popup__overlay -->
      <div class="search-popup__content">
        <form action="#">
          <label for="search" class="sr-only">你想寻找什么？</label
          ><!-- /.sr-only -->
          <input type="text" id="search" placeholder="Search Here..." />
          <button type="submit" aria-label="search submit" class="thm-btn">
            <i class="icon-magnifying-glass"></i>
          </button>
        </form>
      </div>
      <!-- /.search-popup__content -->
    </div>
    <!-- /.search-popup -->

    <a href="#" data-target="html" class="scroll-to-target scroll-to-top"
      ><i class="fa fa-angle-up"></i
    ></a>
  </div>
</template>

<style>
#cnm {
  background-color: #1989fb;
  color: white;
  text-align: center;
  font-size: 150%;
  line-height: 150%;
}
</style>

<script>
import axios from "axios";
const { join } = require("path");
export default {
  head: {
    link: [
      { rel: "preconnect", href: "https://fonts.googleapis.com/" },
      { rel: "preconnect", href: "https://fonts.gstatic.com/" },
      { rel: "stylesheet", href: "css/acs/css2.css" },
      { rel: "stylesheet", href: "css/acs/bootstrap.min.css" },
      { rel: "stylesheet", href: "css/acs/animate.min.css" },
      { rel: "stylesheet", href: "css/acs/custom-animate.css" },
      { rel: "stylesheet", href: "css/acs/all.min.css" },
      { rel: "stylesheet", href: "css/acs/jarallax.css" },
      {
        rel: "stylesheet",
        href: "https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css",
        integrity:
          "sha384-9gVQ4dYFwwWSjIDZnLEWnxCjeSWFphJiwGPXr1jddIhOegiu1FwO5qRGvFXOdJZ4",
        crossorigin: "anonymous",
      },
      { rel: "stylesheet", href: "css/acs/jquery.magnific-popup.css" },
      { rel: "stylesheet", href: "css/acs/nouislider.min.css" },
      { rel: "stylesheet", href: "css/acs/nouislider.pips.css" },
      { rel: "stylesheet", href: "css/acs/odometer.min.css" },
      { rel: "stylesheet", href: "css/acs/swiper.min.css" },
      { rel: "stylesheet", href: "css/acs/style.css" },
      { rel: "stylesheet", href: "css/acs/tiny-slider.min.css" },
      { rel: "stylesheet", href: "css/acs/stylesheet.css" },
      { rel: "stylesheet", href: "css/acs/owl.carousel.min.css" },
      { rel: "stylesheet", href: "css/acs/owl.theme.default.min.css" },
      { rel: "stylesheet", href: "css/acs/twentytwenty.css" },
      { rel: "stylesheet", href: "css/acs/jquery.bxslider.css" },
      { rel: "stylesheet", href: "css/acs/bootstrap-select.min.css" },
      { rel: "stylesheet", href: "css/acs/vegas.min.css" },
      { rel: "stylesheet", href: "css/acs/jquery-ui.css" },
      { rel: "stylesheet", href: "css/acs/timePicker.css" },
      { rel: "stylesheet", href: "css/acs/izeetak.css" },
      { rel: "stylesheet", href: "css/acs/izeetak-responsive.css" },
    ],
    script: [
      {
        src: join("/", `js/ajs/izeetak.js`),
        defer: true,
      },
      {
        src: join("/", `js/ajs/ccc.js`),
        defer: true,
      },
    ],
  },
    async asyncData() {
    // axios.defaults.headers.Cookie = context.req.headers.cookie    
    return axios.get(`/api/getRep`).then((res) => {
      console.log(res.data.list);
      Array.from(res.data.list).forEach(i => {
        i.VotingTime = new Date(parseInt(res.data.list[0].VotingTime) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ')
      })
      return {
        list: res.data.list,
      };
    });
  },
  methods:{
    vote(id){
      return axios
        .get(`/api/voting?Id=` + id)
        .then((res) => {
          if (res.status == 200) {
            alert(res.data.massage)
          }
        });
    }
  }
};
</script>
