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
      <section class="contact-page">
        <div class="container">
          <div class="row">
            <div class="col-xl-4 col-lg-4">
              <div class="contact-page__left">
                <div class="section-title text-left">
                  <span class="section-title__tagline">{{
                    $route.query.ifs
                  }}</span>
                  <h2 class="section-title__title">
                    请输入您今天的绿色生活数量
                  </h2>
                </div>
                <p class="contact-page__text">
                  您将会中和属于您这部分的碳排放量
                </p>
              </div>
            </div>
            <div class="col-xl-8 col-lg-8">
              <div class="contact-page__right">
                <form>
                <div class="row" style="text-align: center">
                  <div class="col-xl-6">
                <div class="mb-3" style="margin-bottom: 3rem!important;">
                <input
                  class="form-control"
                  name="file"
                  type="file"
                  id="formFileMultiple"
                  @change="getFile($event)"
                  multiple=""
                />
              </div>
              </div>
              </div>
                <div class="row" style="text-align: center">
                  <div class="col-xl-6">
                    <div class="comment-form__input-box">
                      <input
                        type="text"
                        style="width: 100%"
                        id="num"
                        name="num"
                        placeholder="数量"
                      />
                    </div>
                  </div>
                </div>
                <div class="row">
                  <div
                    class="col-xl-6"
                    style="margin-bottom: 5%; margin-top: 5%"
                  >
                    <div
                      class="comment-form__input-box"
                      style="font-size: 130%"
                    >
                      <template v-if="$route.query.ifs == '种植植物'">
                        单位:
                        <input
                          id="a"
                          type="text"
                          style="width: 100%"
                          name="company"
                          value="种植植物,颗"
                          readonly
                        />
                      </template>
                      <template v-else-if="$route.query.ifs == '骑电动车'">
                        单位:
                        <input
                          id="a"
                          type="text"
                          style="width: 100%"
                          name="company"
                          value="骑电动车,KM"
                          readonly
                        />
                      </template>
                      <template v-else-if="$route.query.ifs == '徒步出行'">
                        单位:
                        <input
                          id="a"
                          type="text"
                          style="width: 100%"
                          name="company"
                          value="徒步出行,KM"
                          readonly
                        />
                      </template>
                      <template v-else-if="$route.query.ifs == '新能源车'">
                        单位:<input
                          id="a"
                          type="text"
                          style="width: 100%"
                          name="company"
                          value="新能源车,KM"
                          readonly
                        />
                      </template>
                      <template v-else-if="$route.query.ifs == '垃圾回收'">
                        单位:<input
                          id="a"
                          type="text"
                          style="width: 100%"
                          name="company"
                          value="垃圾回收,KG"
                          readonly
                        />
                      </template>
                      <template v-else>
                        单位:<input
                          id="a"
                          type="text"
                          style="width: 100%"
                          name="company"
                          placeholder="请手动输入单位"
                        />
                      </template>
                    </div>
                  </div>
                </div>
                <div class="row">
                  <div class="col-xl-12">
                    <button
                      type="submit"
                      @click="submitForm($event)"
                      style="width: 49%"
                      class="thm-btn comment-form__btn"
                    >
                      确定
                    </button>
                  </div>
                </div>
                </form>
                <div class="result"></div>
                <!-- /.result -->
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
      <!-- /.mobile-nav__content -->
    </div>
    <!-- /.mobile-nav__wrapper -->

    <!-- /.search-popup -->

    <a href="#" data-target="html" class="scroll-to-target scroll-to-top"
      ><i class="fa fa-angle-up"></i
    ></a>
  </div>
</template>

<script>
import axios from "axios";
const { join } = require("path");
export default {
  data() {
    return {
      ifs: this.$route.query.ifs,
    };
  },
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
  methods: {
          getFile(event) {
            this.file = event.target.files[0];
          },
          submitForm(event) {
            let num = document.getElementById("num").value;
            let company = document.getElementById("a").value;
            event.preventDefault();
            let formData = new FormData();
            formData.append('num', num);
            formData.append('company', company);
            formData.append('file', this.file);

            let config = {
              headers: {
                'Content-Type': 'multipart/form-data'
              }
            }

            return axios.post(`/pys/upload`, formData, config).then(function (res) {
              if (res.status === 200) {
                console.log(res);
                return axios.get(
                `/api/load?company=`+
                  res.data.company +
                  "&num=" +
                  res.data.num +
                  "&dw=" +
                  res.data.dw
              ).then((res) => {
                if (res.status == 200) {
                  console.log(res);
                  return location.href=`/index8?ifs=`+
                  res.data.company+
                  `&num=`+ 
                  res.data.num+
                  `&dw=`+ 
                  res.data.dw+
                  `&EN=`+ 
                  res.data.EN+
                  `&massage=`+
                  res.data.massage
                  
                }
              });
              }
            })
          }
        },
  // methods: {
  //   add() {
  //     let num = document.getElementById("num").value;
  //     let company = document.getElementById("a").value;
  //     return axios
  //       .get(
  //         `/api/load?company=`+
  //           this.ifs +
  //           "&num=" +
  //           num +
  //           "&dw=" +
  //           company
  //       )
  //       .then((res) => {
  //         if (res.status == 200) {
  //           this.$router.push({
  //             path: `/index8`,
  //             query: {
  //               ifs: this.ifs,
  //               num: res.data.num,
  //               dw: res.data.dw,
  //               EN: res.data.EN,
  //               massage: res.data.massage,
  //             },
  //           });
  //         }
  //       });
  //   },
  // },
};
</script>
