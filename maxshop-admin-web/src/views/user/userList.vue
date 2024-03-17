<template> 
  <div class="app-container">
    <el-card class="operate-container" shadow="never">
      <i class="el-icon-tickets"></i>
      <span>用户列表</span>
      <!-- <el-button
        class="btn-add"
        @click="handleAddProduct()"
        size="mini">
        添加
      </el-button> -->
    </el-card>
    <div class="table-container">
      <el-table ref="productTable"
                :data="list"
                style="width: 100%"
                @selection-change="handleSelectionChange"
                v-loading="listLoading"
                border>
        <!-- <el-table-column type="selection" align="center"></el-table-column> -->
        <!-- <el-table-column label="编号" align="center">
          <template slot-scope="scope">{{scope.row.id}}</template>
        </el-table-column> -->
        <!-- <el-table-column label="商品图片" width="120" align="center">
          <template slot-scope="scope"><img style="height: 80px" :src="scope.row.pic"></template>
        </el-table-column> -->
        <el-table-column label="用户姓名" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.name}}</p>
          </template>
        </el-table-column>
        <el-table-column label="用户性别" align="center">
          <template slot-scope="scope">
            <p>{{genderDic[scope.row.gender]}}</p>
          </template>
        </el-table-column>
        <el-table-column label="电话" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.mobile}}</p>
          </template>
        </el-table-column>
        <el-table-column label="生日"  align="center">
          <template slot-scope="scope">{{scope.row.birthday}}</template>
        </el-table-column>

      </el-table>
    </div>
    <div class="pagination-container">
      <el-pagination
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        layout="total, sizes,prev, pager, next,jumper"
        :page-size="listQuery.pageSize"
        :page-sizes="[5,10,15]"
        :current-page.sync="pageNum"
        :total="total">
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
  import {getUserList} from '@/apis/goods'
  import {fetchListWithChildren} from '@/api/productCate'
  import {getGoods} from '@/apis/goods'

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
    name: "userList",
    data() {
      return {
        editSkuInfo:{
          dialogVisible:false,
          productId:null,
          productSn:'',
          productAttributeCategoryId:null,
          stockList:[],
          productAttr:[],
          keyword:null,
          pageNum:0
        },
        genderDic: {
          'male': '男',
          'female': '女'
        },
        operates: [
          {
            label: "商品上架",
            value: "publishOn"
          },
          {
            label: "商品下架",
            value: "publishOff"
          },
          {
            label: "设为推荐",
            value: "recommendOn"
          },
          {
            label: "取消推荐",
            value: "recommendOff"
          },
          {
            label: "设为新品",
            value: "newOn"
          },
          {
            label: "取消新品",
            value: "newOff"
          },
          {
            label: "转移到分类",
            value: "transferCategory"
          },
          {
            label: "移入回收站",
            value: "recycle"
          }
        ],
        operateType: null,
        listQuery: {
          pn:1,
          pnum:20
        },
        list: [],
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
      // this.getBrandList();
      // this.getProductCateList();
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
      getList() {
        this.listLoading = true;
        getUserList(this.listQuery).then(response => {
          this.listLoading = false;
          this.list = this.list.concat(response.data);
          this.total = response.total;
          this.pageNum  = this.listQuery.pn
          this.listQuery.pn = response.data.length==this.listQuery.pnum?this.listQuery.pn+2:this.listQuery.pn
        });
      },
      handleSizeChange(val) {
        this.listQuery.pn = 1;
        this.listQuery.pnum = val;
        this.getList();
      },
      handleCurrentChange(val) {
        this.listQuery.pn = val;
        this.getList();
      },
    }
  }
</script>
<style></style>
