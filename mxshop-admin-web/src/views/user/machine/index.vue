<template> 
  <div class="app-container">
   
    
    <div class="table-container">
      <el-table ref="productTable"
                :data="list"
                style="width: 100%"
                v-loading="listLoading"
                border>
                <!-- @selection-change="handleSelectionChange" -->
        <!-- <el-table-column type="selection" align="center"></el-table-column> -->
        <el-table-column label="用户编号" align="center">
          <template slot-scope="scope">{{scope.row.user_id}}</template>
        </el-table-column>
        <el-table-column label="留言" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.message}}</p>
          </template>
        </el-table-column>
        <el-table-column label="留言类型" align="center">
          <template slot-scope="scope">
            <p>{{typeDic[scope.row.type]}}</p>
          </template>
        </el-table-column>
        <el-table-column label="主题" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.subject}}</p>
          </template>
        </el-table-column>
        <el-table-column label="附件" align="center">
          <template slot-scope="scope">
            <!-- <p>{{scope.row.subject}}</p>
             -->
             <a style="color:rgb(64, 158, 255);" :href="scope.row.file" target="_blank">点击下载附件</a>
          </template>
        </el-table-column>
        <!-- <el-table-column label="价格/货号" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.subject}}</p>
          </template>
        </el-table-column> -->
      </el-table>
    </div>
    <div class="pagination-container">
      <el-pagination
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        layout="total, sizes,prev, pager, next,jumper"
        :page-size="listQuery.pnum"
        :page-sizes="[5,10,15]"
        :current-page.sync="pageNum"
        :total="total">
      </el-pagination>
    </div>
  </div>
</template>
<script>
  import {getuserfav,getMessage} from '@/apis/goods'

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
    name: "machineList",
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
        typeDic:{
          1:'留言',
          2:'投诉',
          3:'询问',
          4:'售后',
          5:'求购'
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
      
      }
    },
    created() {
      this.getList();
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
        getMessage().then(response => {
          this.listLoading = false;
          
          this.list =  this.list.concat(response.data);
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