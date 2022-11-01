<template>
  <div class="qrCode">
    <el-row :gutter="20" class="txtElRow">
      <el-col :span="7" class="txtElCol">
        <el-card class="boxCard" :body-style="{height:'100%',width:'100%'}">
          <el-form label-position="top" label-width="100px" :model="formQrSetting">
            <el-form-item label="二维码颜色" ref="formS">
              <el-input style="width: 80%;margin-right: 10px;" v-model="formQrSetting.fontColor"></el-input>
              <el-color-picker @change="generateQRFunc" v-model="formQrSetting.fontColor" color-format="hex" :predefine="predefineColors" />
            </el-form-item>
            <el-form-item label="图片背景色">
              <el-input style="width: 80%;margin-right: 10px;" v-model="formQrSetting.bgColor"></el-input>
              <el-color-picker @change="generateQRFunc" v-model="formQrSetting.bgColor" color-format="hex" :predefine="predefineColors" />
            </el-form-item>
            <el-form-item label="图片尺寸">
              <el-input-number @change="generateQRFunc" controls-position="right" style="width: 100%;" v-model="formQrSetting.qrSize" :min="1" :max="1000" />
            </el-form-item>
            <el-form-item label="图片边框">
              <el-input-number @change="generateQRFunc" controls-position="right" style="width: 100%;" v-model="formQrSetting.qrBorderWidth" :min="0" :max="1000" />
            </el-form-item>
            <el-form-item label="容错率">
              <el-select @change="generateQRFunc" style="width: 100%;" v-model="formQrSetting.qrErrorRate" placeholder="选择容错率">
                <el-option
                    v-for="item in qrErrorRateOptions"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="Logo">
              <el-input readonly style="width: 75%;margin-right: 10px;" v-model="formQrSetting.logoImg" placeholder="">
                <template #append><div @click="openLogoAlertDialog" style="cursor: pointer;"><el-icon><FolderOpened /></el-icon></div></template>
              </el-input>
              <el-button type="info" @click="deleteCommonLogoFunc" icon="Delete" />
            </el-form-item>
            <el-form-item label="Logo边框比例">
              <el-input-number @change="generateQRFunc" controls-position="right" style="width: 100%;" v-model="formQrSetting.logoBorderWidth" :min="0" :precision="2" :step="0.01" :max="1" />
            </el-form-item>
            <el-form-item label="圆角大小">
              <el-input-number @change="generateQRFunc" controls-position="right" style="width: 100%;" v-model="formQrSetting.logoBorderRadius" :min="1" :max="1000" />
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="17" class="txtElCol">
        <el-card class="boxCard" style="height: 40%;" :body-style="{height:'86%',width:'100%'}">
          <el-input
              v-model="formQrSetting.content"
              class="inputData"
              type="textarea"
              placeholder="请输入要生成二维码的文本，可以是： 1、普通文本；2、网址；3、 邮箱；4、以及其他任意文本..."
              resize="none"
              :input-style="{height:'100%',width:'100%'}"
          />
          <div class="line-t-10"></div>
          <div style="width: 100%;height: 50px;">
            <el-button style="float: left;" type="primary" @click="generateQRFunc" icon="Promotion">生成二维码</el-button>
            <el-button style="float: right;" type="danger" @click="saveQrFunc" icon="Download">保存二维码</el-button>
          </div>
        </el-card>
        <div class="line-t-10"></div>
        <el-card class="boxCard" style="height: 58%;" :body-style="{height:'100%',width:'100%'}">
          <div style="height:100%;border: 2px dashed #4d88a5;display: flex;justify-content: center;align-items: center;min-height: 300px; padding: 1rem;background-color: #f6f8f9;overflow-y: scroll;">
<!--            <div v-show="qrBaseData == ''" class="img-placeholder"><span class="placeholder-text">此处将显示二维码图片</span></div>-->
            <div style="max-height: 320px;max-width: 320px;" ref="qrCodeEl"></div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>

</template>

<script>
import { onMounted, reactive, ref} from "vue";
import QrCodeWithLogo from 'qr-code-with-logo'
import {OfficeQrGenerateSave, OpenLogoAlertDialog} from "../../wailsjs/go/main/App.js";
export default {
  name: "QrCode",
  setup(){
    const qrCode = ref()
    const qrCodeEl = ref()
    const isShowDrawer = ref(false)
    const qrBaseData = ref("")
    const formQrSetting = reactive({
      content:"",
      fontColor:"#000000",
      bgColor:"#ffffff",
      qrSize:300,
      qrBorderWidth:1,
      qrErrorRate:"Q",
      logoImg:"",
      logoBorderWidth:0.01,
      logoBorderRadius:1,
      logoFilePath:""
    });
    const qrErrorRateOptions = [
      {value: 'L', label: '7%',},
      {value: "M", label: '15%',},
      {value: "Q", label: '25%',},
      {value: "H", label: '30%',},
    ]
    const deleteCommonLogoFunc = () => {
      formQrSetting.logoImg = "";
      formQrSetting.logoFilePath = "";
      generateQRFunc()
    }
    const openLogoAlertDialog = () => {
      OpenLogoAlertDialog().then( res => {
        formQrSetting.logoImg = res.fileName;
        formQrSetting.logoFilePath = res.filePath;
        generateQRFunc()
      })
    }
    const handleDrawerClose = () => {
      isShowDrawer.value = false;
    }
    const saveQrFunc = () => {
      let base = document.getElementById("myCanvas").toDataURL("image/png")
      base = base.replace("data:image/png;base64,","")
      OfficeQrGenerateSave(base).then( res => {

      })
    }
    
    const generateQRFunc = () => {
      if(formQrSetting.content == ""){
        return;
      }
      let logoOptions = ""
      if(formQrSetting.logoImg != ""){
        logoOptions = {
          src: formQrSetting.logoFilePath,
          logoRadius: formQrSetting.logoBorderRadius,
          borderRadius: formQrSetting.logoBorderRadius,
          borderColor: '#FFFFFF', // IE下 只能使用 6位的 RGB
          borderSize:formQrSetting.logoBorderWidth // 边框大小 相对二维码的比例
        }
      }

      QrCodeWithLogo.toCanvas({
        canvas: document.getElementById("myCanvas"),
        content: formQrSetting.content,
        width: formQrSetting.qrSize,
        nodeQrCodeOptions: {
          margin: formQrSetting.qrBorderWidth,
          errorCorrectionLevel: formQrSetting.qrErrorRate,
          color: {
            dark: formQrSetting.fontColor,
            light: formQrSetting.bgColor
          }
        },
        logo:  logoOptions
      })
    }
    onMounted(()=>{
      const myCanvas = document.createElement('canvas')
      myCanvas.id = "myCanvas"
      myCanvas.style.maxWidth = "320px";
      myCanvas.style.maxHeight = "320px";
      myCanvas.style.cursor = "pointer";
      qrCodeEl.value.appendChild(myCanvas)
    })
    const predefineColors = ref([
      '#ff4500',
      '#ff8c00',
      '#ffd700',
      '#90ee90',
      '#00ced1',
      '#1e90ff',
      '#c71585',
      'rgba(255, 69, 0, 0.68)',
      'rgb(255, 120, 0)',
      'hsv(51, 100, 98)',
      'hsva(120, 40, 94, 0.5)',
      'hsl(181, 100%, 37%)',
      'hsla(209, 100%, 56%, 0.73)',
      '#c7158577',
    ])
    return {
      qrCode,
      qrCodeEl,
      isShowDrawer,
      formQrSetting,
      qrErrorRateOptions,
      predefineColors,
      qrBaseData,
      generateQRFunc,
      deleteCommonLogoFunc,
      handleDrawerClose,
      openLogoAlertDialog,
      saveQrFunc
    }
  }
}
</script>

<style scoped>
.qrCode{
  padding: 20px;
  width: 100%;
  height: 100%;
  position: relative;
}
.qrCode .txtElRow{
  height: 100%;
}
.qrCode .txtElRow .txtElCol .boxCard{
  height: 100%;
}
</style>