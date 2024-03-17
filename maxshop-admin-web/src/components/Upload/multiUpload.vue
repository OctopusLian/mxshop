<template> 
  <div>
    <el-upload
      :action="useOss?host:minioUploadUrl"
      :data="useOss?dataObj:null"
      list-type="picture-card"
      :file-list="value"
      :before-upload="beforeUpload"
      :on-remove="handleRemove"
      :on-success="handleUploadSuccess"
      :on-preview="handlePreview"
      :limit="maxCount"
      :on-exceed="handleExceed"
      :disabled="disabled"
      ref="load"
      :multiple="multiple"
    >
      <!-- :http-request="requestImg" -->
      <i class="el-icon-plus"></i>
    </el-upload>
    <el-dialog :visible.sync="dialogVisible">
      <img width="100%" :src="dialogImageUrl" alt="">
    </el-dialog>
  </div>
</template>
<script>
  import {policy} from '@/api/oss'
  import axios from 'axios'
  import qs from 'qs';
  import send_request from '@/utils/update'
// import { delete } from 'vue/types/umd';

  export default {
    name: 'multiUpload',
    props: {
      //图片属性数组
      value: {
        type: Array,
        default:()=> {
          return []
        }
      },
      multiple: {
        type:Boolean,
        default:true
      },
      //最大上传图片数量
      maxCount:{
        type:Number,
        default:5
      },
      disabled:Boolean
    },
    data() {
      return {
        dataObj: {
          signature: '',
          key: '',
          key: '',
        },
        host:'',
        dir:'',
        dialogVisible: false,
        dialogImageUrl:null,
        useOss:true, //使用oss->true;使用MinIO->false
        ossUploadUrl:'',
        minioUploadUrl:'http://localhost:8080/minio/upload',
      };
    },
    computed: {
      fileList() {

        return this.value||[];
      }
    },
    methods: {
      emitInput(value) {

        this.$emit('input', value)
      },
      handleRemove(file, value) {
        this.emitInput(value);
      },
      handlePreview(file) {
        this.dialogVisible = true;
        this.dialogImageUrl=file.url;
      },
      beforeUpload(file) {
        let _self = this;
          _self.dataObj.name = file.name
        if(!this.useOss){
          //不使用oss不需要获取策略
          return true;
        }
        return new Promise((resolve, reject) => {
        console.log('文件名',file)

        let body = send_request()
            var obj = eval ("(" + body + ")");
            _self.dataObj.policy = obj.policy;
            _self.dataObj.signature = obj.signature;
            _self.dataObj.OSSAccessKeyId = obj.accessid;
            _self.dataObj.key = obj.dir+"${filename}"
            _self.dir = obj.dir
            // _self.dataObj.expire = obj.expire;
            _self.host = obj.host;
            // _self.dataObj.callback = obj.callback;
            _self.dataObj.success_action_status ='200', //让服务端返回200,不然，默认会返回204

            resolve(true)
          }).catch(err => {
            console.log(err)
            reject(false)
          })
        // })
      },
      requestImg(file) {
        console.log(file,axios)
        let params = file.data
        params.file = file.file
        axios.post(this.dataObj.host,qs.stringify(params),{headers:{'Content-Type':'multipart/form-data'}}).then(res=> {

        })
      },
      handleUploadSuccess(res, file,fileList) {
        console.log('图片上传回调', fileList)
        let imgList = []
        fileList.forEach(item=> {
          if(item.size) {
            imgList.push({url:this.host+'/'+this.dir+item.name});
          } else {
            imgList.push(item)
          }
        })
        this.emitInput(imgList);
      },
      handleExceed(files, value) {
        this.$message({
          message: '最多只能上传'+this.maxCount+'张图片',
          type: 'warning',
          duration:1000
        });
      },
    }
  }
</script>
<style>

</style>


