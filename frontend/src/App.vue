<script setup>
import QrCode from "./components/QrCode.vue";</script>

<template>
  <el-container class="elContainer">
    <el-main class="elMain">
      <QrCode></QrCode>
    </el-main>
  </el-container>
</template>

<style>
.elContainer{
  width:100%;
  height:100%;
  position: fixed;
  top: 0px;
  left: 0px;
  overflow: hidden;
}
.elContainer .elMain{
  padding: 0px;
  position: relative;
  overflow: hidden;
}
</style>
