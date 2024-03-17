<template> 
   <div style="margin-top: 50px">
    <el-card shadow="never" class="operate-container">
      <i class="el-icon-tickets"></i>
      <span>轮播图列表</span>
      <el-button size="mini" class="btn-add" @click="handleAdd()">添加</el-button>
    </el-card>
     <div class="table-container">
      <el-table ref="brandTable"
                :data="imgList"
                style="width: 100%"
                border>
                 <el-table-column label="轮播顺序" width="200" align="center">
          <template slot-scope="scope">
            {{scope.row.index}}
          </template>
        </el-table-column>
        <el-table-column label="跳转地址" align="center">
                <!-- @selection-change="handleSelectionChange" -->
          <template slot-scope="scope">
            <a :href="scope.row.url" target="_blank">点击链接跳转</a>
          </template>
                <!-- v-loading="listLoading" -->
        </el-table-column>
        <el-table-column label="图片" align="center">
          <template slot-scope="scope">
            <img :src="scope.row.image" class="imgs" alt="">
          </template>
        </el-table-column>

        <el-table-column label="操作" width="200" align="center">
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="handleUpdate(scope.$index, scope.row)">编辑
            </el-button>
            <el-button
              size="mini"
              type="danger"
              @click="deleteBanner(scope.$index, scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog
      title="添加分类"
      :visible.sync="dialogVisible"
      width="40%">
      <el-form :model="bannerInfo"
               ref="resourceCategoryForm"
               label-width="150px" size="small">
        <el-form-item label="跳转地址：">
          <el-input v-model="bannerInfo.url" style="width: 450px"></el-input>
        </el-form-item>
        <el-form-item label="排序：">
          <el-input v-model="bannerInfo.index" style="width: 450px"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false" size="small">取 消</el-button>
        <el-button type="primary" @click="handleDialogConfirm()" size="small">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="添加分类"
      :visible.sync="dialogVisiblenew"
      width="40%">
      <el-form :model="bannerInfo"
               ref="resourceCategoryForm"
               label-width="150px" size="small">
        <el-form-item label="跳转地址：">
          <el-input v-model="newbannerInfo.url" style="width: 350px"></el-input>
        </el-form-item>
        <el-form-item label="排序：">
          <el-input v-model="newbannerInfo.index" style="width: 350px"></el-input>
        </el-form-item>
        <el-form-item label="图片：">
          <!-- <el-input v-model="newbannerInfo.index" style="width: 450px"></el-input> -->
          <!-- <single-upload v-model="imgs"
                           style="width: 300px;display: inline-block;margin-left: 10px"></single-upload> -->
        <multi-upload v-model="imgs" :disabled='disabledImg' maxCount='1' @input="getImg"></multi-upload>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisiblenew = false" size="small">取 消</el-button>
        <el-button type="primary" @click="finishCommit" size="small">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
  import {fetchListWithChildren} from '@/api/productCate'
  import {fetchList as fetchBrandList} from '@/api/brand'
  import {createBanner,getBanners,deleteBanners,putBanner} from '@/apis/goods';
  import SingleUpload from '@/components/Upload/singleUpload'
  import MultiUpload from '@/components/Upload/multiUpload'

  export default {
    name: "ProductInfoDetail",
    props: {
      // value: Object,
      isEdit: {
        type: Boolean,
        default: false
      }
    },
    components: {SingleUpload,MultiUpload},

    data() {
      return {
        disabled: false,
        isEdit: false,
        imgList:[],
        imgs:[],
        bannerInfo:{},
        newbannerInfo:{},
        disabledImg:false,
        dialogVisiblenew:false,
        dialogVisible:false,
        params:{}
      };
    },
    created() {
      this.getBannerList();
    },
    computed:{
      //商品的编号
      productId(){
        return this.value.id;
      }
    },
    watch: {
    },
    methods: {
      edit() {
        this.disabled = false
      },
      handleAdd() {
        this.dialogVisiblenew = true
      },
      handleUpdate(index,row) {
        this.bannerInfo = row
        this.dialogVisible = true
      },
      handleDialogConfirm() {
         putBanner(this.bannerInfo.id,{index:this.bannerInfo.index,url: this.bannerInfo.url}).then(res=> {
          this.$notify({
                title: '提示',
                message: '修改成功',
                type: 'success'
                });
                this.dialogVisible = false
        }).catch(err=> {
           this.$notify({
                title: '错误',
                message: err.msg,
                type: 'error'
                });
                this.dialogVisible = false
        })
      },
       getImg(data) {
        console.log(data)
        this.newbannerInfo.image = data[0].url

      },
      updateBanner(id,index) {
        this.$prompt('请输入跳转的url','修改', {
          confirmButtonText: '确定',
          cancelButtonText: '取消'
        }).then((res)=> {
          let param = {
            index:index,
            url:res.value
          }
          putBanner(id,{index}).then(res=> {
          this.$notify({
                title: '提示',
                message: '修改成功',
                type: 'success'
                });
        })
        })
      },
      getBannerList() {
        getBanners().then(response => {
          response.forEach(item=> {
            item.url = item.image
          })
          this.imgList = response;
        });
      },
      changeIndex(id,index) {
        putBanner(id,{index}).then(res=> {
          this.$notify({
                title: '提示',
                message: '修改成功',
                type: 'success'
                });
        })
      },
      deleteBanner(id,index) {
        deleteBanners(id).then(res=> {
          this.$notify({
                title: '提示',
                message: '修改成功',
                type: 'success'
                });
                this.imgList.splice(index,1)
        })
      },
      finishCommit(data) {
        // if(this.isEdit) {
          console.log(this.newbannerInfo)
          this.newbannerInfo.index = Number(this.newbannerInfo.index)
          // this.params.image = this.imgs[0].url
          createBanner(this.newbannerInfo).then(res=> {
            this.$notify({
                title: '提示',
                message: '修改成功',
                type: 'success'
                });
                this.dialogVisiblenew = false
                this.newbannerInfo = {}
      this.getBannerList();
          }).catch(err=> {
            this.$notify({
                title: '提示',
                message: '新增错误',
                type: 'error'
                });

                this.dialogVisiblenew = false
                this.newbannerInfo = {}
          })
      },
      handleBrandChange(val) {
        let brandName = '';
        for (let i = 0; i < this.brandOptions.length; i++) {
          if (this.brandOptions[i].value === val) {
            brandName = this.brandOptions[i].label;
            break;
          }
        }
        this.value.brandName = brandName;
      }
    }
  }
</script>

<style scoped lang="scss">
.each-img{
  display: flex;
  width: 500px;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  .img-box{

    width: 300px;
    height: 300px;
    // border: 1px solid;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .buttons{
    width: 150px;
    display: flex;
  }
  .img{
    width: 250px;
    height: 250px;
    // background-color: red;
    border: 1px solid;
  }
}
.imgs{
  width: 50px;
  height: 50px;
}
</style>
