<template> 
  <div>
      <!-- :action="useOss?ossUploadUrl:minioUploadUrl" -->
    <el-upload
      :action="useOss?dataObj.host:minioUploadUrl"

      :data="useOss?dataObj:null"
      list-type="picture"
      :multiple="false" :show-file-list="showFileList"
      :file-list="fileList"
      :before-upload="beforeUpload2"
      :on-remove="handleRemove"
      :on-success="handleUploadSuccess"
      :on-preview="handlePreview">
      <el-button size="small" type="primary">点击上传</el-button>
      <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过10MB</div>
      <!-- <i class="el-icon-plus"></i> -->

    </el-upload>
    <el-dialog :visible.sync="dialogVisible">
      <img width="100%" :src="fileList[0].url" alt="">
    </el-dialog>
  </div>
</template>
<script>
  import {policy} from '@/api/oss'
  import send_request from '@/utils/update'

  export default {
    name: 'singleUpload',
    props: {
      value: String
    },
    computed: {
      imageUrl() {
        return this.value;
      },
      imageName() {
        if (this.value != null && this.value !== '') {
          return this.value.substr(this.value.lastIndexOf("/") + 1);
        } else {
          return null;
        }
      },
      fileList() {
        return [{
          name: this.imageName,
          url: this.imageUrl
        }]
      },
      showFileList: {
        get: function () {
          return this.value !== null && this.value !== ''&& this.value!==undefined;
        },
        set: function (newValue) {
        }
      }
    },
    data() {
      return {
        // dataObj: {
        //   policy: '',
        //   signature: '',
        //   key: '',
        //   ossaccessKeyId: '',
        //   dir: '',
        //   host: '',
        //   // callback:'',
        // },
         dataObj: {
          signature: '',
          key: '',
          key: '',
          host: ''
        },
        dialogVisible: false,
        useOss:true, //使用oss->true;使用MinIO->false
        ossUploadUrl:'http://macro-oss.oss-cn-shenzhen.aliyuncs.com',
        minioUploadUrl:'http://localhost:8080/minio/upload',
      };
    },
    methods: {
      emitInput(val) {
        this.$emit('input', val)
      },
      handleRemove(file, fileList) {
        this.emitInput('');
      },
      handlePreview(file) {
        this.dialogVisible = true;
      },
       beforeUpload2(file) {
        let _self = this;
        if(!this.useOss){
          //不使用oss不需要获取策略
          return true;
        }
        return new Promise((resolve, reject) => {
        let body = send_request()
            var obj = eval ("(" + body + ")");
            _self.dataObj.policy = obj.policy;
            _self.dataObj.signature = obj.signature;
            _self.dataObj.OSSAccessKeyId = obj.accessid;
            _self.dataObj.key = obj.dir
            _self.dataObj.expire = obj.expire;
            _self.dataObj.host = obj.host;
            _self.dataObj.callback = obj.callback;
            _self.dataObj.success_action_status ='200', //让服务端返回200,不然，默认会返回204

            resolve(true)
          }).catch(err => {
            console.log(err)
            reject(false)
          })
        // })
      },
      beforeUpload(file) {
        let _self = this;
        if(!this.useOss){
          //不使用oss不需要获取策略
          return true;
        }
        return new Promise((resolve, reject) => {
          policy().then(response => {
            _self.dataObj.policy = response.data.policy;
            _self.dataObj.signature = response.data.signature;
            _self.dataObj.ossaccessKeyId = response.data.accessKeyId;
            _self.dataObj.key = response.data.dir + '/${filename}';
            _self.dataObj.dir = response.data.dir;
            _self.dataObj.host = response.data.host;
            // _self.dataObj.callback = response.data.callback;
            resolve(true)
          }).catch(err => {
            console.log(err)
            reject(false)
          })
        })
      },
      handleUploadSuccess(res, file) {
        this.showFileList = true;
        this.fileList.pop();
        let url = this.dataObj.host + '/' + this.dataObj.key + file.name;
        if(!this.useOss){
          //不使用oss直接获取图片路径
          url = res.data.url;
        }
        this.fileList.push({name: file.name, url: url});
        this.emitInput(this.fileList[0].url);
      }
    }
  }
</script>
<style>

</style>


