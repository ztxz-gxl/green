<template>
  <div>
    <!-- <div class="preloader">
      <img
        class="preloader__image"
        width="60"
        src="picture/loader.png"
        alt=""
      />
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
                  <a href="javascript:void(0)" class="mobile-nav__toggler"
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
      <!--Projects Page Start-->
      <section class="project-two project-page">
        <div class="container">
          <div class="section-title text-center">
            <span class="section-title__tagline">节能减排的行为</span>
            <h2 class="section-title__title">选择你的今日行为</h2>
          </div>
          <div class="row">
            <div
              class="col-xl-4 col-lg-6 col-md-6 wow fadeInLeft"
              data-wow-delay="100ms"
            >
              <!--Project One Single-->
              <div class="project-one__single">
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '种植植物' },
                  }"
                >
                  <div class="project-one__img">
                    <img src="picture/project-page-img-1.jpg" alt="" />
                  </div>
                  <div class="project-one__content">
                    <p class="project-one__tagline">陶冶情操</p>
                    <h2 class="project-one__title">种植植物</h2>
                    <div class="project-one__arrow">
                      <span class="icon-right-arrow"></span>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
            <div
              class="col-xl-4 col-lg-6 col-md-6 wow fadeInUp"
              data-wow-delay="200ms"
            >
              <!--Project One Single-->
              <div class="project-one__single">
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '骑电动车'},
                  }"
                >
                  <div class="project-one__img">
                    <img src="picture/project-page-img-2.jpg" alt="" />
                  </div>
                  <div class="project-one__content">
                    <p class="project-one__tagline">还能兜风</p>
                    <h2 class="project-one__title">骑电动车</h2>
                    <div class="project-one__arrow">
                      <span class="icon-right-arrow"></span>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
            <div
              class="col-xl-4 col-lg-6 col-md-6 wow fadeInRight"
              data-wow-delay="300ms"
            >
              <!--Project One Single-->
              <div class="project-one__single">
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '徒步出行'},
                  }"
                >
                  <div class="project-one__img">
                    <img src="picture/project-page-img-3.jpg" alt="" />
                  </div>
                  <div class="project-one__content">
                    <p class="project-one__tagline">就是有点累</p>
                    <h2 class="project-one__title">徒步出行</h2>
                    <div class="project-one__arrow">
                      <span class="icon-right-arrow"></span>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
            <div
              class="col-xl-4 col-lg-6 col-md-6 wow fadeInLeft"
              data-wow-delay="400ms"
            >
              <!--Project One Single-->
              <div class="project-one__single">
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '新能源车'},
                  }"
                >
                  <div class="project-one__img">
                    <img src="picture/project-page-img-4.jpg" alt="" />
                  </div>
                  <div class="project-one__content">
                    <p class="project-one__tagline">环保的新能源汽车哦</p>
                    <h2 class="project-one__title">新能源车</h2>
                    <div class="project-one__arrow">
                      <span class="icon-right-arrow"></span>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
            <div
              class="col-xl-4 col-lg-6 col-md-6 wow fadeInUp"
              data-wow-delay="500ms"
            >
              <!--Project One Single-->
              <div class="project-one__single">
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '垃圾回收'},
                  }"
                >
                  <div class="project-one__img">
                    <img src="picture/project-page-img-5.jpg" alt="" />
                  </div>
                  <div class="project-one__content">
                    <p class="project-one__tagline">记得分类</p>
                    <h2 class="project-one__title">垃圾回收</h2>
                    <div class="project-one__arrow">
                      <span class="icon-right-arrow"></span>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
            <div
              class="col-xl-4 col-lg-6 col-md-6 wow fadeInRight"
              data-wow-delay="600ms"
            >
              <!--Project One Single-->
              <div class="project-one__single">
                <n-link
                  :to="{
                    path: '/index5',
                    query: { ifs: '其他行为'},
                  }"
                >
                  <div class="project-one__img">
                    <img src="picture/project-page-img-6.jpg" alt="" />
                  </div>
                  <div class="project-one__content">
                    <p class="project-one__tagline">其他的方法</p>
                    <h2 class="project-one__title">其他行为</h2>
                    <div class="project-one__arrow">
                      <span class="icon-right-arrow"></span>
                    </div>
                  </div>
                </n-link>
              </div>
            </div>
          </div>
        </div>
      </section>
      <!--Projects Page End-->
    </div>
    <!-- /.page-wrapper -->

    <!--手机端导航栏-->
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
    <!-- /.mobile-nav__wrapper -->

    <div class="search-popup">
      <div class="search-popup__overlay search-toggler"></div>
      <!-- /.search-popup__overlay -->
      <div class="search-popup__content">
        <form action="#">
          <label for="search" class="sr-only">search here</label
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

    <a
      href="javascript:void(0)"
      data-target="html"
      class="scroll-to-target scroll-to-top"
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
      {
        src: join("/", `js/ajs/izeetak.js`),
        defer: true,
      },
    ],
  },
};
</script>
