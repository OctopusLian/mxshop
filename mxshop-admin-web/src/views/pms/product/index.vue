<template> 
  <div class="app-container">
    <el-card class="filter-container" shadow="never">
      <div style="margin-bottom:25px">
        <i class="el-icon-search"></i>
        <span>筛选搜索</span>
        <el-button
          style="float: right"
          @click="handleSearchList()"
          type="primary"
          size="small">
          查询结果
        </el-button>
        <el-button
          style="float: right;margin-right: 15px"
          @click="handleResetSearch()"
          size="small">
          重置
        </el-button>
      </div>
      <div style="margin-top: 15px">
        <el-form :inline="true" :model="goodsParams" size="small" label-width="140px">
          <el-form-item label="输入搜索：">
            <el-input style="width: 203px" v-model="goodsParams.q" placeholder="商品名称"></el-input>
          </el-form-item>
          <el-form-item label="商品分类：">
            <el-cascader
              clearable
              v-model="goodsParams.c"
              @change="getBrand"
              :props="{value:'id',label:'name',children:'sub_category'}"
              :options="productCateOptions">
            </el-cascader>
          </el-form-item>
          <el-form-item label="商品品牌：">
            <el-select v-model="goodsParams.b" placeholder="请选择品牌" clearable>
              <el-option
                v-for="item in brandOptions"
                :key="item.id"
                :label="item.name"
                :value="item.id">
              </el-option>
            </el-select>
          </el-form-item>
        </el-form>
      </div>
    </el-card>
    <el-card class="operate-container" shadow="never">
      <i class="el-icon-tickets"></i>
      <span>数据列表</span>
      <el-button
        class="btn-add"
        @click="handleAddProduct()"
        size="mini">
        添加
      </el-button>
    </el-card>
    <div class="table-container">
      <el-table ref="productTable"
                :data="list"
                style="width: 100%"
                @selection-change="handleSelectionChange"
                v-loading="listLoading"
                border>
        <!-- <el-table-column type="selection" align="center"></el-table-column> -->
        <el-table-column label="编号" align="center" width="80">
          <template slot-scope="scope">{{scope.$index+1}}</template>
        </el-table-column>
        <el-table-column label="商品名称" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.name}}</p>
          </template>
        </el-table-column>
        <el-table-column label="商品品牌" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.brand.name}}</p>
          </template>
        </el-table-column>
        <el-table-column label="商品分类" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.category.name}}</p>
          </template>
        </el-table-column>
        <el-table-column label="价格/货号" width="120" align="center">
          <template slot-scope="scope">
            <p>价格：￥{{scope.row.shop_price}}</p>
            <!-- <p>货号：{{scope.row.goods_brief}}</p> -->
          </template>
        </el-table-column>
        <el-table-column label="标签" align="center">
          <template slot-scope="scope">
            <p>上架：
              <el-switch
                @change="handlePublishStatusChange('sale',scope.$index, scope.row)"
                :active-value="true"
                :inactive-value="false"
                v-model="scope.row.on_sale">
              </el-switch>
            </p>
            <p>新品：
              <el-switch
                @change="handlePublishStatusChange('new',scope.$index, scope.row)"
                :active-value="true"
                :inactive-value="false"
                v-model="scope.row.is_new">
              </el-switch>
            </p>
            <p>推荐：
              <el-switch
                @change="handlePublishStatusChange('hot',scope.$index, scope.row)"
                :active-value="true"
                :inactive-value="false"
                v-model="scope.row.is_hot">
              </el-switch>
            </p>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="160" align="center">
          <template slot-scope="scope">
            <p>
              <el-button
                size="mini"
                @click="handleShowProduct(scope.$index, scope.row)">查看详情
              </el-button>
              <!-- <el-button
                size="mini"
                @click="handleUpdateProduct(scope.$index, scope.row)">编辑
              </el-button> -->
            </p>
            <p>
              <!-- <el-button
                size="mini"
                @click="handleShowLog(scope.$index, scope.row)">日志
              </el-button> -->
              <el-button
                size="mini"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)">删除
              </el-button>
            </p>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="pagination-container">
      <el-pagination
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        layout="total, sizes,prev, pager, next,jumper"
        :page-size="goodsParams.pnum"
        :current-page.sync="goodsParams.pn"
        :total="total">
        <!-- :page-sizes="[5,10,15]" -->
      </el-pagination>
    </div>

  </div>
</template>
<script>
  import {
    fetchList,
    updateDeleteStatus,
    updateNewStatus,
    updateRecommendStatus,
    updatePublishStatus
  } from '@/api/product'
  import {fetchList as fetchSkuStockList,update as updateSkuStockList} from '@/api/skuStock'
  import {fetchList as fetchProductAttrList} from '@/api/productAttr'
  import {fetchList as fetchBrandList} from '@/api/brand'
  import {fetchListWithChildren} from '@/api/productCate'
  import {getGoods,deleteGoods,getBrands,getCategorys,putGoodsStatus,getBrandsByCate } from '@/apis/goods'

  const defaultListQuery = {
    keyword: null,
    pageNum: 1,
    pageSize: 5,
    publishStatus: null,
    verifyStatus: null,
    productSn: null,
    productCategoryId: null,
    brandId: null
  };
  export default {
    name: "productList",
    data() {
      return {
        editSkuInfo:{
          dialogVisible:false,
          productId:null,
          productSn:'',
          productAttributeCategoryId:null,
          stockList:[],
          productAttr:[],
          keyword:null
        },
        goodsParams:{
          pn:1,
          pnum:20,
          c:''
        },
        operates: [
          {
            label: "商品上架",
            value: "publishOn"
          }, {
            label: "新品",
            value: "publishOn"
          }, {
            label: "推荐",
            value: "publishOn"
          }, {
            label: "删除",
            value: "delete"
          },
        ],
        operateType: null,
        listQuery: Object.assign({}, defaultListQuery),
        list: null,
        total: null,
        listLoading: true,
        selectProductCateValue: null,
        multipleSelection: [],
        productCateOptions: [],
        brandOptions: [],
        publishStatusOptions: [{
          value: 1,
          label: '上架'
        }, {
          value: 0,
          label: '下架'
        }],
        verifyStatusOptions: [{
          value: 1,
          label: '审核通过'
        }, {
          value: 0,
          label: '未审核'
        }]
      }
    },
    created() {
      this.getList();
      this.getBrandList();
      this.getProductCateList();
    },
    watch: {
      selectProductCateValue: function (newValue) {
        if (newValue != null && newValue.length == 2) {
          this.listQuery.productCategoryId = newValue[1];
        } else {
          this.listQuery.productCategoryId = null;
        }

      }
    },
    filters: {
      verifyStatusFilter(value) {
        if (value === 1) {
          return '审核通过';
        } else {
          return '未审核';
        }
      }
    },
    methods: {
      getProductSkuSp(row, index) {
        let spData = JSON.parse(row.spData);
        if(spData!=null&&index<spData.length){
          return spData[index].value;
        }else{
          return null;
        }
      },
      getBrand(id) {
        console.log(id)
        this.goodsParams.c = id[2]
        getBrandsByCate(id[2]).then(res=> {
          this.brandOptions=res
        })
      },
      getList() {
        this.listLoading = true;
          console.log(this.goodsParams)
        getGoods(this.goodsParams).then(response => {
          this.listLoading = false;
          console.log("goods_list")
          console.log(response.data)
          this.list = response.data;
          this.total = response.total;
        });
      },
      getBrandList() {
        getBrands().then(response => {
          this.brandOptions = response;
          let brandList = response;
          // for (let i = 0; i < brandList.length; i++) {
          //   this.brandOptions.push({label: brandList[i].name, value: brandList[i].id});
          // }
        });
      },
      getProductCateList() {
        getCategorys().then(response => {
          let list = response;
          this.productCateOptions = list;
          // for (let i = 0; i < list.length; i++) {
          //   let children = [];
          //   if (list[i].children != null && list[i].children.length > 0) {
          //     for (let j = 0; j < list[i].children.length; j++) {
          //       children.push({label: list[i].children[j].name, value: list[i].children[j].id});
          //     }
          //   }
          //   this.productCateOptions.push({label: list[i].name, value: list[i].id, children: children});
          // }
        });
      },

      handleSearchEditSku(){
        fetchSkuStockList(this.editSkuInfo.productId,{keyword:this.editSkuInfo.keyword}).then(response=>{
          this.editSkuInfo.stockList=response.data;
        });
      },

      handleSearchList() {
        this.listQuery.pageNum = 1;
        this.getList();
      },
      handleAddProduct() {
        this.$router.push({path:'/addProduct'});
      },
      handleBatchOperate() {
        if(this.operateType==null){
          this.$message({
            message: '请选择操作类型',
            type: 'warning',
            duration: 1000
          });
          return;
        }
        if(this.multipleSelection==null||this.multipleSelection.length<1){
          this.$message({
            message: '请选择要操作的商品',
            type: 'warning',
            duration: 1000
          });
          return;
        }
        this.$confirm('是否要进行该批量操作?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          let ids=[];
          for(let i=0;i<this.multipleSelection.length;i++){
            ids.push(this.multipleSelection[i].id);
          }
          switch (this.operateType) {
            case this.operates[0].value:
              this.updatePublishStatus(true,ids);
              break;
            case this.operates[1].value:
              this.updatePublishStatus(false,ids);
              break;
            case this.operates[2].value:
              this.updateRecommendStatus(1,ids);
              break;
            case this.operates[3].value:
              this.updateRecommendStatus(0,ids);
              break;
            case this.operates[4].value:
              this.updateNewStatus(1,ids);
              break;
            case this.operates[5].value:
              this.updateNewStatus(0,ids);
              break;
            case this.operates[6].value:
              break;
            case this.operates[7].value:
              this.updateDeleteStatus(1,ids);
              break;
            default:
              break;
          }
          this.getList();
        });
      },
      handleSizeChange(val) {
        this.goodsParams.pn = 1;
        this.goodsParams.pnum = val;
        this.getList();
      },
      handleCurrentChange(val) {
        this.goodsParams.pn = val;
        this.getList();
      },
      handleSelectionChange(val) {
        this.multipleSelection = val;
      },
      handlePublishStatusChange(paramname,index, row) {
          if (index===0) {
              this.updatePublishStatus(paramname,false, row);
          }else if(index===1){
              this.updatePublishStatus(paramname,true, row);
          }

      },
      handleNewStatusChange(index, row) {
        let ids = [];
        ids.push(row.id);
        this.updateNewStatus(row.newStatus, ids);
      },
      handleRecommendStatusChange(index, row) {
        let ids = [];
        ids.push(row.id);
        this.updateRecommendStatus(row.recommandStatus, ids);
      },
      handleResetSearch() {
        this.selectProductCateValue = [];
        this.listQuery = Object.assign({}, defaultListQuery);
      },
      handleDelete(index, row){
        this.$confirm('是否要进行删除商品?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          deleteGoods(row.id).then(res=> {
            this.$message({
            message: '删除成功',
            type: 'success',
            duration: 1000
          });
          this.list.splice(index,1)
          });
        });
      },
      handleUpdateProduct(index,row){
        this.$router.push({path:'/updateProduct',query:{id:row.id}});
      },
      handleShowProduct(index,row){
        console.log("handleShowProduct",row);
        this.$router.push({path:'/updateProduct',query:{id:row.id}});
      },
      handleShowVerifyDetail(index,row){
        console.log("handleShowVerifyDetail",row);
      },
      handleShowLog(index,row){
        console.log("handleShowLog",row);
      },
      updatePublishStatus(paramname,param, row) {
        console.log('数据',row)
        let params = {
          'sale':row.on_sale,
          'hot':row.is_hot,
          'new':row.is_new

        }
        params[paramname] = param
        putGoodsStatus(row.id,params).then(response => {
          this.$message({
            message: '修改成功',
            type: 'success',
            duration: 1000
          });
        }).catch(err=> {
          console.log(err)
           this.$message({
            message: err.msg,
            type: 'success',
            duration: 1000
          });
        });
      },
      updateNewStatus(newStatus, ids) {
        let params = new URLSearchParams();
        params.append('ids', ids);
        params.append('newStatus', newStatus);
        updateNewStatus(params).then(response => {
          this.$message({
            message: '修改成功',
            type: 'success',
            duration: 1000
          });
        });
      },
      updateRecommendStatus(recommendStatus, ids) {
        let params = new URLSearchParams();
        params.append('ids', ids);
        params.append('recommendStatus', recommendStatus);
        updateRecommendStatus(params).then(response => {
          this.$message({
            message: '修改成功',
            type: 'success',
            duration: 1000
          });
        });
      },
      updateDeleteStatus(deleteStatus, ids) {
        let params = new URLSearchParams();
        params.append('ids', ids);
        params.append('deleteStatus', deleteStatus);
        updateDeleteStatus(params).then(response => {
          this.$message({
            message: '删除成功',
            type: 'success',
            duration: 1000
          });
        });
        this.getList();
      }
    }
  }
</script>
<style></style>


