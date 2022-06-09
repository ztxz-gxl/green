<template>
  <div>
    <div class="page-wrapper">
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
                  <span class="section-title__tagline">{{ ifs }}</span>
                  <h2 class="section-title__title">今日行为添加成功！</h2>
                </div>
                <p class="contact-page__text">将会给您中和您所需中和的碳排放</p>
              </div>
            </div>
          </div>
        </div>
        <div class="container">
          <div class="row">
            <div class="col-xl-4 col-lg-4" style="width:240px;">
              <div class="contact-details__single">
                <div class="contact-details__icon">
                  <span class="icon-map"></span>
                </div>
                <div class="contact-details__content">
                  <p class="contact-details__sub-title">{{ ifs }}</p>
                  <h2>{{ num }}</h2>
                </div>
              </div>
            </div>

            <div class="col-xl-4 col-lg-4" style="width:240px;">
              <div class="contact-details__single ">
                <div class="contact-details__icon">
                  <span class="icon-map"></span>
                </div>
                <div class="contact-details__content">
                  <p class="contact-details__sub-title">中和能</p>
                  <h2>{{ num }} {{ dw }}</h2>
                </div>
              </div>
            </div>

            <div class="col-xl-4 col-lg-4" style="width:240px;">
              <div class="contact-details__single ">
                <div class="contact-details__icon">
                  <span class="icon-map"></span>
                </div>
                <div class="contact-details__content">
                  <p class="contact-details__sub-title">获得奖励</p>
                  <h2>{{ EN }} GTH</h2>
                </div>
              </div>
            </div>
            <div class="col-xl-4 col-lg-4" style="width:240px">
              <div class="contact-details__single ">
                <div class="contact-details__icon">
                  <span class="icon-map"></span>
                </div>
                <div class="contact-details__content">
                  <p class="contact-details__sub-title">交易ID</p>
                  <h2>ID:{{ massage }}</h2>
                </div>
              </div>
            </div>
          </div>
          <HR />
          <div class="row">
            <div class="col-xl-12">
              <button
                class="thm-btn comment-form__btn"
                @click="QD"
                >确定</button
              >
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
            :to="{ path: '/index2' }"
            aria-label="logo image"
            ><img src="" width="155" alt=""
          /></n-link>
        </div>
        <!-- /.logo-box -->
        <div class="mobile-nav__container"></div>
        <!-- /.mobile-nav__container -->
      </div>
      <!-- /.mobile-nav__content -->
      <a href="#" data-target="html" class="scroll-to-target scroll-to-top"
        ><i class="fa fa-angle-up"></i
      ></a>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      ifs: this.$route.query.ifs,
      num: this.$route.query.num,
      dw: this.$route.query.dw,
      EN: this.$route.query.EN,
      massage: this.$route.query.massage,
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
  },
   methods: {
    QD() {
            this.$router.push({
              path: `/index2`,
            });
          }
  },
};
</script>
