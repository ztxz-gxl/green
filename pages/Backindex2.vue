<template>
  <div>
    <div class="lyear-layout-web">
      <div class="lyear-layout-container">
            <nav class="navbar navbar-default">
          <div class="topbar">
            <div class="topbar-left">
              <!-- <a href="index.html"><img src="images/logo-sidebar.png" title="LightYear" alt="LightYear" /></a> -->
              <n-link :to="{ name: 'Backindex' }" style="color: black"
                ><span style="font-size: 20px" class="navbar-page-title">
                  后台数据中心
                </span></n-link
              >
            </div>

            <ul class="topbar-right">
              <li class="dropdown dropdown-profile" style="display: flex">
                <a href="javascript:void(0)" data-toggle="dropdown">
                  <img
                    class="img-avatar img-avatar-48 m-r-10"
                    src="images/users/avatar.jpg"
                    alt="绿色生活"
                  />
                  <span style="font-size: 20px"
                    >绿色生活 <span class="caret"></span
                  ></span>
                </a>
                <ul class="dropdown-menu dropdown-menu-right">
                  <li>
                    <n-link :to="{ name: 'Backindex1' }"
                      ><i class="mdi mdi-account"></i> 查看举报信息</n-link
                    >
                  </li>
                  <li>
                    <n-link :to="{ name: 'Backindex4' }"
                      ><i class="mdi mdi-lock-outline"></i>提取银行代币</n-link
                    >
                  </li>
                  <li>
                    <n-link :to="{ name: 'BackLogin',params:{id:this,id} }"
                      ><i class="mdi mdi-lock-outline"></i>退出</n-link
                    >
                  </li>
                </ul>
              </li>
            </ul>
          </div>
        </nav>
        <div class="container-fluid">
          <div class="row">
            <div class="col-lg-12">
              <div class="card">
                <div class="card-header">
                  <h4>举报信息</h4>
                </div>
                <div class="card-body">
                  <div class="table-responsive">
                    <table style="font-size: 16px" class="table table-hover">
                      <thead>
                        <tr>
                          <th>序号</th>
                          <th>用户地址</th>
                          <th>日期</th>
                          <th>用户姓名</th>
                          <th>状态</th>
                        </tr>
                      </thead>
                      <tbody>
                        <tr v-for="i in list" :key="i.Id">
                          <td>{{ i.Id }}</td>
                          <td>{{ i.Addr }}</td>
                          <td>{{ i.Date }}</td>
                          <td>{{ i.Fun }}</td>
                          <td>
                            <span class="label label-success">{{ i.Stu }}</span>
                          </td>
                        </tr>
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
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
  data() {
    return {
      id: this.$route.params.id,
    };
  },
  head: {
    link: [
      { rel: "stylesheet", href: "css/bootstrap.min.css" },
      { rel: "stylesheet", href: "css/materialdesignicons.min.css" },
      { rel: "stylesheet", href: "css/style.min.css" },
    ],
  },
  asyncData() {
    return axios.get(`/api/fy`).then((res) => {
      return {
        list: res.data.list,
      };
    });
  },
};
</script>
