<template>
  <div class="app-container">
    <el-card class="operate-container" shadow="never">
      <i class="el-icon-tickets" style="margin-top: 5px"></i>
      <span style="margin-top: 5px">数据列表</span>
      <el-button
        class="btn-add"
        @click="goAdd()"
        size="mini">
        添加
      </el-button>
    </el-card>
    <div class="table-container">
        <el-table ref="productTable"
                :data="list"
                style="width: 100%"
                v-loading="listLoading"
                border>
        <el-table-column label="分类" align="center">
          <template slot-scope="scope">{{scope.row.category.name}}</template>
        </el-table-column>
        <el-table-column label="品牌" align="center">
          <template slot-scope="scope">
            <p>{{scope.row.brand.name}}</p>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="goDetail(scope.$index, scope.row.category.id,scope.row)"
            >查看</el-button>
            <el-button
              size="mini"
              @click="deleteBrandToCate(scope.$index,scope.row.id)"
            >删除</el-button>
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
        :page-size="listQuery.pnum"
        :current-page.sync="pageNum"
        :total="total">
        <!-- :page-sizes="[5,10,15]" -->
      </el-pagination>
    </div>
  </div>
</template>

<script>
  import {getBrandToCate,deleteCategorys,putCategorys,deleteBrandToCate} from '@/apis/goods'
import { getToken, setToken, removeToken } from '@/utils/auth'

  export default {
    name: "brandToCate",
    data() {
      return {
        list: [],
        total: null,
        listLoading: true,
        listQuery: {
          pn: 1,
          pnum: 5
        },
        parentId: 0
      }
    },
    created() {
      // this.resetParentId();
      
      this.getList();
    },
    watch: {
      // $route(route) {
      //   // this.resetParentId();
      //   this.getList();
      // }
    },
    methods: {
      getList() {
        this.listLoading = true;
        getBrandToCate().then(response => {
          console.log(response)
          this.listLoading = false;
          
          this.list = this.list.concat(response.data||[]);
          this.total = response.total;
          this.totalPage = response.data.totalPage;
          this.pageSize = response.data.pageSize;
          this.pageNum  = this.listQuery.pn
          this.listQuery.pn = response.data.length==this.listQuery.pnum?this.listQuery.pn+2:this.listQuery.pn
          // this.total = response;
        });
      },
      deleteBrandToCate(index,id) {
        this.$confirm('是否要确认删除?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                  deleteBrandToCate(id).then(response => {
                    this.$message({
                      message: '修改成功！',
                      type: 'success'
                    });
                    this.editDialogVisible =false;
                    this.list.splice(index,1)
                    // this.getList();
                  })
              })
      },
      goAdd() {
        this.$router.push('/editbrandcate')
      },
      goDetail(index,id,obj) {
        this.$router.push({path:'/editbrandcate',query:{id:id}})
        localStorage.setItem('brandToCate',JSON.stringify(obj))
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

<style scoped>
.cat-name{
  margin-right: 50px;
}
</style>
