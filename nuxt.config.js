export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'nuxt-test',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      // { hid: 'description', name: 'description', content: '' },
      // { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
    ],
    script:[
      { src: 'js/jquery.min.js', defer: true },
      { src: 'js/bootstrap.min.js', defer: true },
      { src: 'js/perfect-scrollbar.min.js', defer: true },
      { src: 'js/main.min.js', defer: true },
      { src:'js/ajs/jquery-3.6.0.min.js', defer: true },
      { src:'js/ajs/bootstrap.bundle.min.js', defer: true },
      { src:'js/ajs/jarallax.min.js', defer: true },
      { src:'js/ajs/jquery.ajaxchimp.min.js', defer: true },
      { src:'js/ajs/jquery.appear.min.js', defer: true },
      { src:'js/ajs/jquery.circle-progress.min.js', defer: true },
      { src:'js/ajs/jquery.validate.min.js', defer: true },
      { src:'js/ajs/nouislider.min.js', defer: true },
      { src:'js/ajs/odometer.min.js', defer: true },
      { src:'js/ajs/swiper.min.js', defer: true },
      { src:'js/ajs/tiny-slider.min.js', defer: true },
      { src:'js/ajs/wNumb.min.js', defer: true },
      { src:'js/ajs/wow.js', defer: true },
      { src:'js/ajs/isotope.js', defer: true },
      { src:'js/ajs/countdown.min.js', defer: true },
      { src:'js/ajs/owl.carousel.min.js', defer: true },
      { src:'js/ajs/twentytwenty.js', defer: true },
      { src:'js/ajs/jquery.event.move.js', defer: true },
      { src:'js/ajs/jquery.bxslider.min.js', defer: true },
      { src:'js/ajs/bootstrap-select.min.js', defer: true },
      { src:'js/ajs/vegas.min.js', defer: true },
      { src:'js/ajs/jquery-ui.js', defer: true },
      { src:'js/Chart.js', defer: true },
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/axios','@nuxtjs/proxy'
  ],
  axios: {
      proxy: true, // ??????????????????
      prefix: '/api', 
      prefix: '/pys',// ???????????????url???????????? /api
      credentials: true // ?????????????????????????????????????????????
  },
  proxy: {
    '/api': {
      target: 'http://127.0.0.1:8080', // ??????????????????
      changeOrigin: true, // ??????????????????
      pathRewrite: {
        '^/api': '/', // ??? /api ????????? /
      }
    },
    '/pys': {
      target: 'http://127.0.0.1:5000', // ??????????????????
      changeOrigin: true, // ??????????????????
      pathRewrite: {
        '^/pys': '/', // ??? /api ????????? /
      }
    }
  },
  build: {
    vendor: ['axios'] //?????????????????????
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },
}
