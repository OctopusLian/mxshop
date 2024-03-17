<template> 
  <div class="app-container">
    
    <el-card class="operate-container" shadow="never">
      <i class="el-icon-tickets"></i>
      <span>数据列表</span>
    </el-card>
    <div class="table-container">
      <el-table ref="orderTable"
                :data="list"
                style="width: 100%;"
                @selection-change="handleSelectionChange"
                v-loading="listLoading" border>
        <!-- <el-table-column type="selection" width="60" align="center"></el-table-column> -->
        <el-table-column label="编号" width="80" align="center">
          <template slot-scope="scope">{{scope.row.id}}</template>
        </el-table-column>
        <el-table-column label="收货地址" width="180" align="center">
          <template slot-scope="scope">{{scope.row.address}}</template>
        </el-table-column>
        <el-table-column label="联系电话" width="180" align="center">
          <!-- <template slot-scope="scope">{{scope.row.mobile | formatCreateTime}}</template> -->
          <template slot-scope="scope">{{scope.row.mobile}}</template>
        </el-table-column>
        <el-table-column label="联系人" align="center">
          <template slot-scope="scope">{{scope.row.name}}</template>
        </el-table-column>
        <el-table-column label="订单金额" width="120" align="center">
          <template slot-scope="scope">￥{{scope.row.total}}</template>
        </el-table-column>
        <el-table-column label="支付方式" width="120" align="center">
          <template slot-scope="scope">{{scope.row.post | formatPayType}}</template>
        </el-table-column>
        <!-- <el-table-column label="订单来源" width="120" align="center">
          <template slot-scope="scope">{{scope.row.sourceType | formatSourceType}}</template>
        </el-table-column> -->
        <el-table-column label="订单状态" width="120" align="center">
          <template slot-scope="scope">{{scope.row.status | formatStatus}}</template>
        </el-table-column>
        <el-table-column label="操作" width="200" align="center">
          <template slot-scope="scope">
            <el-button
              size="mini"
              @click="handleViewOrder(scope.$index, scope.row)"
            >查看订单</el-button>
            <!-- <el-button
              size="mini"
              @click="handleCloseOrder(scope.$index, scope.row)"
              v-show="scope.row.status===0">关闭订单</el-button>
            <el-button
              size="mini"
              @click="handleDeliveryOrder(scope.$index, scope.row)"
              v-show="scope.row.status===1">订单发货</el-button>
            <el-button
              size="mini"
              @click="handleViewLogistics(scope.$index, scope.row)"
              v-show="scope.row.status===2||scope.row.status===3">订单跟踪</el-button>
            <el-button
              size="mini"
              type="danger"
              @click="handleDeleteOrder(scope.$index, scope.row)"
              v-show="scope.row.status===4">删除订单</el-button> -->
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
        :current-page.sync="pageNum"
        :page-size="listQuery.pnum"
        :page-sizes="[5,10,15]"
        :total="total">
      </el-pagination>
    </div>
    <el-dialog
      title="关闭订单"
      :visible.sync="closeOrder.dialogVisible" width="30%">
      <span style="vertical-align: top">操作备注：</span>
      <el-input
        style="width: 80%"
        type="textarea"
        :rows="5"
        placeholder="请输入内容"
        v-model="closeOrder.content">
      </el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="closeOrder.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleCloseOrderConfirm">确 定</el-button>
      </span>
    </el-dialog>
    <logistics-dialog v-model="logisticsDialogVisible"></logistics-dialog>
  </div>
</template>
<script>
  import {fetchList,closeOrder,deleteOrder} from '@/api/order'
  import {getOrder} from '@/apis/goods'
  import {formatDate} from '@/utils/date';
  import LogisticsDialog from '@/views/oms/order/components/logisticsDialog';
  const defaultListQuery = {
    pageNum: 1,
    pageSize: 10,
    orderSn: null,
    receiverKeyword: null,
    status: null,
    orderType: null,
    sourceType: null,
    createTime: null,
  };
  export default {
    name: "orderList",
    components:{LogisticsDialog},
    data() {
      return {
        listQuery: Object.assign({}, defaultListQuery),
        listLoading: true,
        list: [],
        total: null,
        operateType: null,
        pageNum:0,
        multipleSelection: [],
        closeOrder:{
          dialogVisible:false,
          content:null,
          orderIds:[]
        },
        statusOptions: [
          {
            label: '待付款',
            value: 0
          },
          {
            label: '待发货',
            value: 1
          },
          {
            label: '已发货',
            value: 2
          },
          {
            label: '已完成',
            value: 3
          },
          {
            label: '已关闭',
            value: 4
          }
        ],
        orderTypeOptions: [
          {
            label: '正常订单',
            value: 0
          },
          {
            label: '秒杀订单',
            value: 1
          }
        ],
        sourceTypeOptions: [
          {
            label: 'PC订单',
            value: 0
          },
          {
            label: 'APP订单',
            value: 1
          }
        ],
        operateOptions: [
          {
            label: "批量发货",
            value: 1
          },
          {
            label: "关闭订单",
            value: 2
          },
          {
            label: "删除订单",
            value: 3
          }
        ],
        logisticsDialogVisible:false
      }
    },
    created() {
      this.getList();
    },
    filters: {
      formatCreateTime(time) {
        let date = new Date(time);
        return formatDate(date, 'yyyy-MM-dd hh:mm:ss')
      },
      formatPayType(value) {
        if (value === 1) {
          return '支付宝';
        } else if (value === 2) {
          return '微信';
        } else {
          return '未支付';
        }
      },
      formatSourceType(value) {
        if (value === 1) {
          return 'APP订单';
        } else {
          return 'PC订单';
        }
      },
      formatStatus(value) {
        if (value === 1) {
          return '待发货';
        } else if (value === 2) {
          return '已发货';
        } else if (value === 3) {
          return '已完成';
        } else if (value === 4) {
          return '已关闭';
        } else if (value === 5) {
          return '无效订单';
        } else {
          return '待付款';
        }
      },
    },
    methods: {
      handleResetSearch() {
        this.listQuery = Object.assign({}, defaultListQuery);
      },
      handleSearchList() {
        this.listQuery.pageNum = 1;
        this.getList();
      },
      handleSelectionChange(val){
        this.multipleSelection = val;
      },
      handleViewOrder(index, row){
        this.$router.push({path:'/oms/orderDetail',query:{id:row.id}})
      },
      handleCloseOrder(index, row){
        this.closeOrder.dialogVisible=true;
        this.closeOrder.orderIds=[row.id];
      },
      handleDeliveryOrder(index, row){
        let listItem = this.covertOrder(row);
        this.$router.push({path:'/oms/deliverOrderList',query:{list:[listItem]}})
      },
      handleViewLogistics(index, row){
        this.logisticsDialogVisible=true;
      },
      handleDeleteOrder(index, row){
        let ids=[];
        ids.push(row.id);
        this.deleteOrder(ids);
      },
      handleBatchOperate(){
        if(this.multipleSelection==null||this.multipleSelection.length<1){
          this.$message({
            message: '请选择要操作的订单',
            type: 'warning',
            duration: 1000
          });
          return;
        }
        if(this.operateType===1){
          //批量发货
          let list=[];
          for(let i=0;i<this.multipleSelection.length;i++){
            if(this.multipleSelection[i].status===1){
              list.push(this.covertOrder(this.multipleSelection[i]));
            }
          }
          if(list.length===0){
            this.$message({
              message: '选中订单中没有可以发货的订单',
              type: 'warning',
              duration: 1000
            });
            return;
          }
          this.$router.push({path:'/oms/deliverOrderList',query:{list:list}})
        }else if(this.operateType===2){
          //关闭订单
          this.closeOrder.orderIds=[];
          for(let i=0;i<this.multipleSelection.length;i++){
            this.closeOrder.orderIds.push(this.multipleSelection[i].id);
          }
          this.closeOrder.dialogVisible=true;
        }else if(this.operateType===3){
          //删除订单
          let ids=[];
          for(let i=0;i<this.multipleSelection.length;i++){
            ids.push(this.multipleSelection[i].id);
          }
          this.deleteOrder(ids);
        }
      },
      handleSizeChange(val){
        this.listQuery.pn = 1;
        this.listQuery.pnum = val;
        this.getList();
      },
      handleCurrentChange(val){
        this.listQuery.pn = val;
        this.getList();
      },
      handleCloseOrderConfirm() {
        if (this.closeOrder.content == null || this.closeOrder.content === '') {
          this.$message({
            message: '操作备注不能为空',
            type: 'warning',
            duration: 1000
          });
          return;
        }
        let params = new URLSearchParams();
        params.append('ids', this.closeOrder.orderIds);
        params.append('note', this.closeOrder.content);
        closeOrder(params).then(response=>{
          this.closeOrder.orderIds=[];
          this.closeOrder.dialogVisible=false;
          this.getList();
          this.$message({
            message: '修改成功',
            type: 'success',
            duration: 1000
          });
        });
      },
      getList() {
        this.listLoading = true;
        getOrder(this.listQuery).then(response => {
          this.listLoading = false;
          this.list = this.list.concat(response.data);
          this.total = response.total;
          this.pageNum = this.listQuery.pn
          this.listQuery.pn = response.data.length==this.listQuery.pnum?this.listQuery.pn+1:this.listQuery.pn

        });
      },
      deleteOrder(ids){
        this.$confirm('是否要进行该删除操作?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          let params = new URLSearchParams();
          params.append("ids",ids);
          deleteOrder(params).then(response=>{
            this.$message({
              message: '删除成功！',
              type: 'success',
              duration: 1000
            });
            this.getList();
          });
        })
      },
      covertOrder(order){
        let address=order.receiverProvince+order.receiverCity+order.receiverRegion+order.receiverDetailAddress;
        let listItem={
          orderId:order.id,
          orderSn:order.orderSn,
          receiverName:order.receiverName,
          receiverPhone:order.receiverPhone,
          receiverPostCode:order.receiverPostCode,
          address:address,
          deliveryCompany:null,
          deliverySn:null
        };
        return listItem;
      }
    }
  }
</script>
<style scoped>
  .input-width {
    width: 203px;
  }
</style>


