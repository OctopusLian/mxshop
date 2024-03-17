<template> 
    <div class="brand-cate">
        <el-form :model="productCateInfo"
             ref="productCateFrom"
             label-width="150px">

            <el-form-item label="分类：" prop="name">
                <el-cascader
                    v-model="productCateInfo.category_id"
                    :options="productCateOptions"
                    :props="props"></el-cascader>
        </el-form-item>

            <el-form-item label="品牌：">
                <el-select v-model="productCateInfo.brand_id"
                        placeholder="请选择品牌" >
                <el-option
                    v-for="item in brandOptions"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit('productCateFrom')">提交</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
import {getBrands,getCategorys,getBrandToCateDetail,createBrandToCate,putBrandToCate} from '@/apis/goods'
  export default {
    name: 'editBrandToCate',
    data() {
        return{
            productCateInfo:{},
            props:{
                children:'sub_category',
                label:'name',
                value:'id'
            },
            productCateOptions:[],
            brandOptions:[]
        }
    },
    created() {
        if(this.$route.query.id) {
            // getBrandToCateDetail(this.$route.query.id).then(res=> {
                let obj = JSON.parse(localStorage.getItem('brandToCate'))||{}
                this.productCateInfo.category_id = obj.category.id
                this.productCateInfo.brand_id = obj.brand.id
                console.log(this.productCateInfo)
                // this.productCateInfo.id = obj.id
                // localStorage.delete('brandToCate')
            // })
        }

      this.getBrandList();
      this.getProductCateList();
    },
    methods:{
        getBrandList() {
        getBrands().then(response => {
          this.brandOptions = response.data;
        });
      },
      getProductCateList() {
        getCategorys().then(response => {
          this.productCateOptions = response;
          let _this = this
          if(this.productCateInfo.category_id){
              response.forEach(res=> {
                  if(res.sub_category&&res.sub_category.length>0) {
                    res.sub_category.forEach(item=> {
                        if(item.sub_category&&item.sub_category.length>0) {
                            item.sub_category.forEach(sub=> {
                                if(sub.id ==_this.productCateInfo.category_id){
                                    _this.productCateInfo.category_id=[res.id,item.id,sub.id]
                                }
                            })

                        }
                    })
                  }
              })
          }
        });
      },
      onSubmit() {
          console.log(this.productCateInfo)
          this.productCateInfo.category_id = this.productCateInfo.category_id[2]
          if(this.$route.query.id) {
              putBrandToCate(this.$route.query.id,this.productCateInfo).then(res=> {
                    this.$message({
                        type: 'success',
                        message: '修改成功!'
                    });
                    this.$router.push('/brandcate')
              }).catch(err=> {
                  console.log(err)
                  this.$message({
                        type: 'success',
                        message: err.msg
                    });
              })
          } else {
              createBrandToCate(this.productCateInfo).then(res=> {
                    this.$message({
                        type: 'success',
                        message: '新增成功!'
                    });
                    this.$router.push('/brandcate')
              }).catch(err=> {
                  console.log(err)
                  this.$message({
                        type: 'success',
                        message: err.smg
                    });
              })
          }
      }
    }
  }
</script>
<style>
.brand-cate{
    margin-top: 30px
}
</style>
