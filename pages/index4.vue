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
        <!-- /.sticky-header__content -->
      </div>
      <!-- /.stricky-header -->

      <!--Page Header Start-->
      <section class="page-header">
        <div
          class="page-header-bg"
          style="background-image: url(image/page-header-bg.jpg)"
        ></div>
        <div class="page-header-bg-overly"></div>
        <div
          class="page-header-shape wow slideInRight"
          data-wow-delay="100ms"
          data-wow-duration="2500ms"
          style="background-image: url(image/page-header-shape.png)"
        ></div>
        <div class="container">
          <div class="page-header__inner">
            <ul class="thm-breadcrumb list-unstyled">
              <li>
                <n-link :to="{ path: '/' }"
                  >首页</n-link
                >
              </li>
              <li><span>/</span></li>
              <li>基金</li>
            </ul>
            <h2>绿色银行</h2>
          </div>
        </div>
      </section>
      <!--Page Header End-->

      <!--Help Start-->
      <section class="help">
        <div class="container">
          <div class="row">
            <div class="col-xl-12">
              <ul class="help__box list-unstyled clearfix">
                <li
                  class="help__single help__box-one wow fadeInLeft"
                  data-wow-delay="100ms"
                >
                  <div class="help__box-one-content">
                    <h3 class="help__box-one-title">我们的基金可以做什么？</h3>
                  </div>
                  <div class="help__box-one-img">
                    <img src="picture/help-box-one-img-1.jpg" alt="" />
                  </div>
                </li>
                <li
                  class="help__single help__box-two wow fadeInLeft"
                  data-wow-delay="300ms"
                >
                  <div class="help__box-two-content">
                    <div class="help__box-two-icon">
                      <span class="icon-learning"></span>
                    </div>
                    <h3 class="help__box-two-title">
                      <a href="javascript:void(0)">存储代币</a>
                    </h3>
                    <p class="help__box-two-text">将奖励的储存起来</p>
                  </div>
                </li>
                <li
                  class="help__single help__box-two wow fadeInLeft"
                  data-wow-delay="600ms"
                >
                  <div class="help__box-two-content">
                    <div class="help__box-two-icon">
                      <span class="icon-verify"></span>
                    </div>
                    <h3 class="help__box-two-title">
                      <a href="javascript:void(0)">代币激励</a>
                    </h3>
                    <p class="help__box-two-text">
                      储存的代币将会获得一定的代币利息
                    </p>
                  </div>
                </li>
                <li
                  class="
                    help__single help__box-two help__box-last
                    wow
                    fadeInLeft
                  "
                  data-wow-delay="900ms"
                >
                  <div class="help__box-two-content">
                    <div class="help__box-two-icon">
                      <span class="icon-cloud"></span>
                    </div>
                    <h3 class="help__box-two-title">
                      <a href="javascript:void(0)">绿色公益</a>
                    </h3>
                    <p class="help__box-two-text">绿色银行将会将您的代币用来</p>
                  </div>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>
      <section class="contact-page">
        <div class="container">
          <div class="row" style="border: solid; border-width: 1px;border-radius: 5px;">
            <div id="cnm">
              <p>利息公布(每半年刷新一次)</p>
            </div>
            <table style="text-align: center">
              <thead style="background: aquamarine; color: black">
                <tr>
                  <th>序号</th>
                  <th>利率</th>
                  <th>存储时间</th>
                </tr>
              </thead>
              <tbody style="background-color: aliceblue">
                <tr v-for="(i, index) in list" :key="index">
                  <td>{{ index + 1 }}</td>
                  <td>{{ i.Text }}%</td>
                  <td>{{ i.Type }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </section>
      <section class="contact-page">
        <div class="container">
          <div class="row">
            <div class="col-xl-4 col-lg-4">
              <div class="contact-page__left">
                <div class="section-title text-left">
                  <span class="section-title__tagline">绿色银行</span>
                  <h2 class="section-title__title">选择要储存的能量数量</h2>
                </div>
                <p class="contact-page__text">
                  通过我们的智能合约将能量储存兑换成代币奖励
                </p>
              </div>
            </div>
            <div class="col-xl-8 col-lg-8">
              <div class="contact-page__right">
                <n-link :to="{ path: '/index6' }">
                  <div class="row" style="margin-top: 10%; margin-left: 25%">
                    <div class="col-xl-12">
                      <button type="submit" class="thm-btn comment-form__btn">
                        加入绿色银行
                      </button>
                    </div>
                    </div>
                    </n-link>
                    <n-link :to="{ path: '/index6' }">
                    <div class="row" style="margin-top: 10%; margin-left: 25%">
                    <div class="col-xl-12" style="margin-top: 10%;">
                      <button type="submit" class="thm-btn comment-form__btn">
                        取出银行代币
                      </button>
                    </div>
                  </div>
                </n-link>
                <!-- </form> -->
                <div class="result"></div>
                <!-- /.result -->
              </div>
            </div>
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
  asyncData() {
    return axios.get(`/api/getBankData`).then((res) => {
      return {
        list: res.data.listData4,
      };
    });
  },
};
</script>
