<template>
  <el-card class="form-container" shadow="never">
    <el-form :model="productCateInfo"
             :rules="rules"
             ref="productCateFrom"
             label-width="150px">
      <el-form-item label="分类名称：" prop="name">
        <el-input v-model="productCateInfo.name"></el-input>
      </el-form-item>
      <el-form-item label="上级分类：">
        <el-select v-model="productCateInfo.parent_first"
                   placeholder="请选择第一级分类" @change="chooseCateFirst()">
          <el-option
            v-for="item in productCate"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select>
        <el-select v-if="selectProductCateList.length>0" v-model="productCateInfo.parent_second"
                   placeholder="请选择二级分类">
          <el-option
            v-for="item in selectProductCateList"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select>
        <!-- <el-select v-model="productCate.parentId"
                   placeholder="请选择二级分类">
          <el-option
            v-for="item in selectProductSecondCateList"
            :key="item.id"
            :label="item.name"
            :value="item.id">
          </el-option>
        </el-select> -->
      </el-form-item>
      <!-- <el-form-item label="排序：">
        <el-input v-model="productCateInfo.sort"></el-input>
      </el-form-item> -->
      <el-form-item label="是否显示在导航栏：">
        <el-radio-group v-model="productCateInfo.is_tab">
          <el-radio :label="true">是</el-radio>
          <el-radio :label="false">否</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit('productCateFrom')">提交</el-button>
        <el-button v-if="!isEdit" @click="resetForm('productCateFrom')">重置</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>

<script>
  import {postCategorys,putCategorys,getCategoryDetail,getCategorys} from '@/apis/goods';
  import SingleUpload from '@/components/Upload/singleUpload';
// import { delete } from 'vue/types/umd';

  const defaultProductCate = {
    is_tab: '',
    icon: '',
    keywords: '',
    name: '',
    navStatus: 0,
    parentId: 0,
    productUnit: '',
    showStatus: 0,
    productAttributeIdList: []
  };
  export default {
    name: "ProductCateDetail",
    components: {SingleUpload},
    props: {
      isEdit: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        productCate: {},
        selectProductCateList: [],
        selectProductCateList:[],
        productCateInfo:{},
        defaultCate: [
          {id: 0, name: '无此级分类'}
        ],
        rules: {
          name: [
            {required: true, message: '请输入品牌名称', trigger: 'blur'},
          ],
          is_tab: [
            {required: true, message: '请选择是否显示在导航栏', trigger: 'blur'},
          ],
        },
        filterAttrs: [],
        filterProductAttrList: [{
          value: []
        }]
      }
    },
    created() {
      getCategorys().then(res=> {
        this.productCate =this.defaultCate.concat(res)
      })
      if (this.isEdit) {
        getCategoryDetail(this.$route.query.id).then(response => {
          this.productCateInfo = response;
        });
      } else {
        // this.productCate = Object.assign({}, defaultProductCate);
      }
    },
    methods: {
      getProductAttrCateList() {
        fetchListWithAttr().then(response => {
          let list = response.data;
          for (let i = 0; i < list.length; i++) {
            let productAttrCate = list[i];
            let children = [];
            if (productAttrCate.productAttributeList != null && productAttrCate.productAttributeList.length > 0) {
              for (let j = 0; j < productAttrCate.productAttributeList.length; j++) {
                children.push({
                  label: productAttrCate.productAttributeList[j].name,
                  value: productAttrCate.productAttributeList[j].id
                })
              }
            }
            this.filterAttrs.push({label: productAttrCate.name, value: productAttrCate.id, children: children});
          }
        });
      },
      chooseCateFirst() {
        console.log(this.productCateInfo.parent_first)
        this.productCate.forEach(item=> {
          console.log(this.productCateInfo.parent_first,item)
          if(this.productCateInfo.parent_first==item.id) {
  
            this.selectProductCateList = this.defaultCate.concat(item.sub_category)
          }
        })
        console.log(this.selectProductCateList)
      },
      getProductAttributeIdList() {
        //获取选中的筛选商品属性
        let productAttributeIdList = [];
        for (let i = 0; i < this.filterProductAttrList.length; i++) {
          let item = this.filterProductAttrList[i];
          if (item.value !== null && item.value.length === 2) {
            productAttributeIdList.push(item.value[1]);
          }
        }
        return productAttributeIdList;
      },
      onSubmit(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$confirm('是否提交数据', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              if(this.productCateInfo.parent_second) {
                this.productCateInfo.level=3
                this.productCateInfo.parent=this.productCateInfo.parent_second
                delete this.productCateInfo.parent_second
              } else if(this.productCateInfo.parent_first&&!this.productCateInfo.parent_second){
                this.productCateInfo.parent=this.productCateInfo.parent_first
                delete this.productCateInfo.parent_first
                this.productCateInfo.level=2
              } else{
                this.productCateInfo.level=1
              }
              let {name,parent,level,is_tab} = this.productCateInfo
              let params = {
                name,parent,level,is_tab
              }
              if (this.isEdit) {
                this.productCate.productAttributeIdList = this.getProductAttributeIdList();
                putCategorys(this.$route.query.id, params).then(response => {
                  this.$message({
                    message: '修改成功',
                    type: 'success',
                    duration: 1000
                  });
                  this.$router.back();
                });
              } else {
                // this.productCate.productAttributeIdList = this.getProductAttributeIdList();
                postCategorys(params).then(response => {
                  // this.$refs[formName].resetFields();
                  // this.resetForm(formName);
                  this.$message({
                    message: '提交成功',
                    type: 'success',
                    duration: 1000
                  });
                  this.$router.push('/productCate')

                });
              }
            });

          } else {
            this.$message({
              message: '验证失败',
              type: 'error',
              duration: 1000
            });
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
        this.productCate = Object.assign({}, defaultProductCate);
        this.getSelectProductCateList();
        this.filterProductAttrList = [{
          value: []
        }];
      },
      removeFilterAttr(productAttributeId) {
        if (this.filterProductAttrList.length === 1) {
          this.$message({
            message: '至少要留一个',
            type: 'warning',
            duration: 1000
          });
          return;
        }
        var index = this.filterProductAttrList.indexOf(productAttributeId);
        if (index !== -1) {
          this.filterProductAttrList.splice(index, 1)
        }
      },
      handleAddFilterAttr() {
        if (this.filterProductAttrList.length === 3) {
          this.$message({
            message: '最多添加三个',
            type: 'warning',
            duration: 1000
          });
          return;
        }
        this.filterProductAttrList.push({
          value: null,
          key: Date.now()
        });
      }
    },
    filters: {
      filterLabelFilter(index) {
        if (index === 0) {
          return '筛选属性：';
        } else {
          return '';
        }
      }
    }
  }
</script>

<style scoped>

</style>
