<template>
  <div class="container">
    <div class="form-warp">
      <form class="sign-in-form">
        <h2 class="form-title">登录</h2>
        <input placeholder="用户名" v-model="userLoginForm.username" />
        <input
          type="password"
          placeholder="密码"
          v-model="userLoginForm.password"
        />
        <div class="submit-btn" @click="handleSubmit">立即登录</div>
      </form>
      <form class="sign-up-form">
        <h2 class="form-title">注册</h2>
        <input placeholder="用户名" />
        <input type="password" placeholder="密码" />
        <div class="submit-btn">立即注册</div>
      </form>
    </div>
    <div class="desc-warp">
      <div class="desc-warp-item sign-up-desc">
        <div class="content">
          <button id="sign-up-btn">注册</button>
        </div>
        <img src="../static/img/svg/log.svg" alt="" />
      </div>
      <div class="desc-warp-item sign-in-desc">
        <div class="content">
          <button id="sign-in-btn">登录</button>
        </div>
        <img src="../static/img/svg/register.svg" alt="" />
      </div>
    </div>
  </div>
</template>

<script>
import { HttpManager } from "../api/index";

export default {
  name: "Login",
  data() {
    return {
      userLoginForm: {
        username: "admin",
        password: "123456",
      },
    };
  },
  mounted() {
    const singUpBtn = document.querySelector("#sign-up-btn");
    const singInBtn = document.querySelector("#sign-in-btn");
    const container = document.querySelector(".container");

    singUpBtn.addEventListener("click", () => {
      container.classList.add("sign-up-mode");
    });
    singInBtn.addEventListener("click", () => {
      container.classList.remove("sign-up-mode");
    });
  },
  methods: {
    // 登陆提交
    handleSubmit() {
      let params = {
        adminName: this.userLoginForm.username,
        adminPassword: this.userLoginForm.password,
      };
      HttpManager.create(params)
        .then((res) => {
          console.log(res);
        })
        .catch();
    },
  },
};
</script>

<style  scoped>
@import url("../static/css/login.css");
</style>
