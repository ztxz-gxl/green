<template>
  <div>
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
                        <li>
                          <n-link
                            :to="{ path: '/approve' }"
                            >代币授权</n-link
                          >
                        </li>
                        <li>
                          <n-link
                            :to="{ path: '/transfer' }"
                            >代币转账</n-link
                          >
                        </li>
                        <li>
                          <n-link
                            :to="{ path: '/info' }"
                            >代币信息</n-link
                          >
                        </li>
                      </ul>
                    </li>
                    <li>
                      <n-link :to="{ path: '/my' }"
                        >个人中心</n-link
                      >
                    </li>
                    <li>
                      <n-link :to="{ path: '/index7' }">切换账户</n-link>
                    </li>
                    <li>
                      <n-link :to="{ path: '/index7' }">退出</n-link>
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
      <!--Contact Page Start-->
      <section class="contact-page">
        <div class="container">
          <div class="row">
            <div class="col-xl-4 col-lg-4">
              <div class="contact-page__left">
                <div class="section-title text-left">
                  <span class="section-title__tagline">兑换货币</span>
                  <h2 class="section-title__title">请输入兑换的货币数量</h2>
                </div>
                <p class="contact-page__text">通过智能合约为您提供服务</p>
              </div>
            </div>
            <div class="col-xl-8 col-lg-8">
              <div class="contact-page__right">
                <form
                  action="assets/inc/sendemail.php"
                  class="comment-one__form contact-form-validated"
                  novalidate="novalidate"
                >
                  <div class="row">
                    <div class="col-xl-6">
                      <div class="comment-form__input-box">
                        <input type="text" id="dhAmonut" placeholder="兑换数量" />
                      </div>
                    </div>
                  </div>

                  <div class="row">
                    <div class="col-xl-4">
                      <a href="javascript:void(0)" @click="DH" class="myButton">交易</a>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </section>
      <section class="contact-page">
        <div class="container">
          <div class="row">
            <div class="col-xl-4 col-lg-4">
              <div class="contact-page__left">
                <div class="section-title text-left">
                  <span class="section-title__tagline">授权转账</span>
                  <h2 class="section-title__title">
                    请输入授权地址接受地址和授权数量
                  </h2>
                </div>
                <p class="contact-page__text">通过智能合约为您提供授权服务</p>
              </div>
            </div>
            <div class="col-xl-8 col-lg-8">
              <div class="contact-page__right">
                <form
                  action="assets/inc/sendemail.php"
                  class="comment-one__form contact-form-validated"
                  novalidate="novalidate"
                >
                  <div class="row">
                    <div class="col-xl-6">
                      <div class="comment-form__input-box">
                        <input type="text" id="spender" placeholder="授权地址" />
                      </div>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-xl-6">
                      <div class="comment-form__input-box">
                        <input type="text" id="addr" placeholder="接受地址" />
                      </div>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-xl-6">
                      <div class="comment-form__input-box">
                        <input type="text" id="amonut" placeholder="授权数量" />
                      </div>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-xl-6">
                      <a href="javascript:void(0)" @click="SQZZ" class="myButton">交易</a>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </section>
      <section class="contact-page">
        <div class="container">
          <div class="row">
            <div class="col-xl-4 col-lg-4">
              <div class="contact-page__left">
                <div class="section-title text-left">
                  <span class="section-title__tagline">转账</span>
                  <h2 class="section-title__title">请输入接受地址和转账数量</h2>
                </div>
                <p class="contact-page__text">通过智能合约为您提供转账服务</p>
              </div>
            </div>
            <div class="col-xl-8 col-lg-8">
              <div class="contact-page__right">
                <form
                  action="assets/inc/sendemail.php"
                  class="comment-one__form contact-form-validated"
                  novalidate="novalidate"
                >
                  <div class="row">
                    <div class="col-xl-6">
                      <div class="comment-form__input-box">
                        <input type="text" id="zzAddr" placeholder="接受地址" />
                      </div>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-xl-6">
                      <div class="comment-form__input-box">
                        <input type="text" id="zzAmount" placeholder="转账数量" />
                      </div>
                    </div>
                  </div>
                  <div class="row">
                    <div class="col-xl-4">
                      <a href="javascript:void(0)" @click="ZZ" class="myButton">交易</a>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </section>
      <!--Contact Page End-->
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
    </div>
    <!-- /.mobile-nav__content -->
    <a href="#" data-target="html" class="scroll-to-target scroll-to-top"
      ><i class="fa fa-angle-up"></i
    ></a>
  </div>
</template>

<style>
.myButton {
  background-color: #007bff;
  border-radius: 20px;
  display: inline-block;
  cursor: pointer;
  color: #ffffff;
  font-family: Arial;
  font-size: 18px;
  padding: 16px 33px;
  text-decoration: none;
  text-shadow: 0px 1px 0px #ffffff;
}
.myButton:hover {
  background-color: #007bff;
}
.myButton:active {
  position: relative;
  top: 1px;
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
  methods:{
    DH(){
      let amount = document.getElementById('dhDmount').value
      return axios
        .get(`/api/DH?amount=` + amount)
        .then((res) => {
          if (res.status == 200) {
            alert(res.data.massage)
          }
        });
    },
    SQZZ(){
      let addr = document.getElementById('addr').value
      let amount = document.getElementById('amount').value
      let spender = document.getElementById('spender').value
      return axios
        .get(`/api/SQZZ?addr=`+ addr +"&spender=" + spender  + "&amount=" + amount)
        .then((res) => {
          if (res.status == 200) {
            alert(res.data.massage)
          }
        });
    },
    ZZ(){
      let addr = document.getElementById('zzAddr').value
      let amount = document.getElementById('zzAmount').value
      return axios
        .get(`/api/ZZ?addr=` + addr + "&amount=" + amount)
        .then((res) => {
          if (res.status == 200) {
            alert(res.data.massage)
          }
        });
    },
  }
};
</script>