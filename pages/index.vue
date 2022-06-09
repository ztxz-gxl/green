<template>
  <div>
    <div class="page-wrapper">
      <header class="main-header main-header-two clearfix">
        <nav class="main-menu main-menu-two clearfix">
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

      <div class="stricky-header stricked-menu main-menu main-menu-two">
        <div class="sticky-header__content"></div>
        <!-- /.sticky-header__content -->
      </div>
      <!-- /.stricky-header -->

      <!--Main Slider Two Start-->
      <section class="main-slider main-slider-two">
        <div class="swiper-container thm-swiper__slider">
          <div class="swiper-wrapper">
            <div class="swiper-slide">
              <div
                class="image-layer"
                style="background-image: url(image/main-slider-2-2.jpg)"
              ></div>
              <div class="image-layer-overlay"></div>
              <div class="container">
                <div class="row">
                  <div class="col-lg-12">
                      <div class="main-slider-two__map"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
      <!--Main Slider Two End-->

      <!--Business Growth Start-->
      <section class="business-growth">
        <div class="business-growth__top">
          <div class="business-growth-bg">
            <div id="particles-js"></div>
          </div>
          <div class="container">
            <div class="row">
              <div class="col-xl-6 col-lg-6">
                <div class="business-growth__left">
                  <div class="section-title text-left">
                    <span class="section-title__tagline">碳中和</span>
                    <h2 class="section-title__title">
                      实现二氧化碳在人类社会与自然环境内的产销平衡
                    </h2>
                  </div>
                </div>
              </div>
              <div class="col-xl-6 col-lg-6">
                <div class="business-growth__right">
                  <p class="business-growth__right-text">
                    碳中和是碳排放组织测算在一个时期内直接或间接产生的碳排放总
                    量，通过其他减碳行为，诸如使用绿色能源、节能减排、植树造林等抵
                    消自身产生的二氧化碳排放量，达到二氧化碳相对“零排放”。目前，碳
                    中和是阻止温室气体影响气候变化的重要目标和方向。

                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="container">
          <div class="business-growth__bottom">
            <div class="row">
              <div
                class="col-xl-4 col-lg-4 wow fadeInLeft"
                data-wow-delay="100ms"
              >
                <!--Business Growth Single-->
                <n-link :to="{ path: '/index2' }">
                  <div class="business-growth__single">
                    <div class="business-growth__content">
                      <h4 class="business-growth__title">绿色出行</h4>
                      <p class="business-growth__text">步行出门或者自行车</p>
                    </div>
                  </div>
                </n-link>
              </div>
              <div
                class="col-xl-4 col-lg-4 wow fadeInUp"
                data-wow-delay="200ms"
              >
                <!--Business Growth Single-->
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '垃圾回收' },
                  }"
                >
                  <div class="business-growth__single">
                    <!-- <div class="business-growth__img">
                                    <img src="picture/business-growth-img-2.jpg" alt="">
                                </div> -->
                    <div class="business-growth__content">
                      <h4 class="business-growth__title">垃圾回收</h4>
                      <p class="business-growth__text">
                        垃圾分类或者减少一次性餐具
                      </p>
                    </div>
                  </div>
                </n-link>
              </div>
              <div
                class="col-xl-4 col-lg-4 wow fadeInRight"
                data-wow-delay="300ms"
              >
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '种植植物'},
                  }"
                >
                  <div class="business-growth__single">
                    <div class="business-growth__content">
                      <h4 class="business-growth__title">种植植物</h4>
                      <p class="business-growth__text">
                        种植植物是实现碳中和的最好方法
                      </p>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
          </div>
        </div>
      </section>
      <!--Business Growth End-->

      <!--Welcome One Start-->
      <section class="welcome-one">
        <div class="container">
          <div class="row">
            <div class="col-xl-6">
              <div
                class="welcome-one__left wow slideInLeft"
                data-wow-delay="100ms"
                data-wow-duration="2500ms"
              >
                <div class="welcome-one__img-box">
                  <div class="welcome-one__img">
                    <img src="picture/welcome-one-img-1.jpg" alt="" />
                  </div>
                  <div class="welcome-one__certified">
                    <div class="welcome-one__certified-icon">
                      <span class="icon-customer"></span>
                    </div>
                    <h4 class="welcome-one__certified-text">
                      JL拍<br />
                      的樱花
                    </h4>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-xl-6">
              <div class="welcome-one__right">
                <div class="section-title text-left">
                  <span class="section-title__tagline">我们可以提供什么</span>
                  <h2 class="section-title__title">
                    加入我们，一起碳中和
                  </h2>
                </div>
                <div class="welcome-one__funded">
                  <div class="welcome-one__funded-img">
                    <img src="" alt="" />
                  </div>
                  <div class="welcome-one__funded-content">
                    <h3 class="welcome-one__funded-text">
                      上传今日绿色行为，来获得节约的
                      <span>中和量</span>
                    </h3>
                  </div>
                </div>
                <div class="welcome-one__funded">
                  <div class="welcome-one__funded-img">
                    <img src="" alt="" />
                  </div>
                  <div class="welcome-one__funded-content">
                    <h3 class="welcome-one__funded-text">
                      储存超出的绿色中和量，将会兑换成
                      <span>代币奖励</span>
                    </h3>
                  </div>
                </div>
                <div class="welcome-one__funded">
                  <div class="welcome-one__funded-img">
                    <img src="" alt="" />
                  </div>
                  <div class="welcome-one__funded-content">
                    <h3 class="welcome-one__funded-text">
                      了解个人数据，并且可以查看今日榜单
                      <span>数据</span>
                    </h3>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
      <!--Welcome One End-->
    </div>
    <!-- /.page-wrapper -->

    <!--侧边导航栏-->
    <div class="mobile-nav__wrapper">
      <div class="mobile-nav__overlay mobile-nav__toggler"></div>
      <!-- /.mobile-nav__overlay -->
      <div class="mobile-nav__content">
        <span class="mobile-nav__close mobile-nav__toggler"
          ><i class="fa fa-times"></i
        ></span>
        <div class="logo-box">
          <n-link
            :to="{ path: '/' }"
            aria-label="logo image"
            ><img src="" width="155" alt=""
          /></n-link>
        </div>
        <!-- /.logo-box -->
        <div class="mobile-nav__container"></div>
        <!-- /.mobile-nav__container -->
      </div>
      <!-- /.mobile-nav__content -->
    </div>

    <a href="#" data-target="html" class="scroll-to-target scroll-to-top"
      ><i class="fa fa-angle-up"></i
    ></a>
  </div>
</template>
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
      { src: join("/", `js/ajs/timePicker.js`), defer: true },
      {
        src: join("/", `js/ajs/izeetak.js`),
        defer: true,
      },
    ],
  },
  data() {
    return {
      id: this.$route.query.id,
    };
  },
};
</script>
