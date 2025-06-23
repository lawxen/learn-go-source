<template>
    <div class="container">
        <hr>
        <form role="form">
            <div class="form-group">
                <label class="text-info text-center" for="q">
                    <h4>查询条件</h4>
                </label>
                <!-- v-model是创建双向数据绑定-->
                <input type="text" id="q" class="form-control"
                       placeholder="输入产品名" v-model="q">
            </div>
            <div>
                <!--@click指定触发的函数，即绑定事件-->
                <input type="button" value="查询"
                       class="btn btn-primary" @click="add()">
            </div>
        </form>
        <hr>
        <div class="text-info text-center">
            <h2>产品信息表</h2>
        </div>
        <table class="table table-bordered table-hover">
            <tr class="text-danger text-center">
                <th>序号</th>
                <th>产品</th>
                <th>数量</th>
                <th>类型</th>
                <th>操作</th>
            </tr>
            <!--遍历输出vue定义的数组-->
            <tr class="text-center" v-for="(item,index) in myData" :key="index">
                <th>{{index+1}}</th>
                <th>{{item.name}}</th>
                <th>{{item.quantity}}</th>
                <th>{{item.kinds}}</th>
                <th>
                    <!--data-target指向模态框-->
                    <!--为每个按钮设置变量nowIndex，用于识别行数-->
                    <button data-toggle="modal" class="btn btn-primary btn-sm"
                            @click="nowIndex=index,message=0" data-target="#layer"
                    >删除
                    </button>
                </th>
            </tr>
            <tr v-show="myData.length!==0">
                <td colspan="5" class="text-right">
                    <!--变量nowIndex设为-2，在deleteMsg函数清空数组myData-->
                    <button data-toggle="modal" class="btn btn-danger btn-sm"
                            @click="nowIndex=-2,message=-1" data-target="#layer">
                        删除全部
                    </button>
                </td>
            </tr>
            <tr v-show="myData.length===0">
                <td colspan="5" class="text-center text-muted">
                    <p>暂无数据....</p>
                </td>
            </tr>
        </table>
        <!--模态框（提示框）-->
        <div role="dialog" class="modal fade bs-example-modal-sm" id="layer">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <!--判断message，选择删除提示语-->
                        <h4 class="modal-title" v-if="message===0">删除吗</h4>
                        <h4 class="modal-title" v-else>删除全部吗</h4>
                        <button type="button" class="close" data-dismiss="modal">
                            <span>&times;</span>
                        </button>
                    </div>
                    <div class="modal-body text-right">
                        <!--触发删除函数deleteMsg-->
                        <button data-dismiss="modal"
                                class="btn btn-primary btn-sm">取消
                        </button>
                        <button data-dismiss="modal" class="btn btn-danger btn-sm"
                                @click="deleteMsg(nowIndex)">确认
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Product',
        data() {
            return {
                q: '',
                myData: [],
                nowIndex: -100,
                message: 0
            }
        },
        methods: {
            // 定义add函数，访问后台获取数据并写入数组myData
            add: function () {
                this.axios.get('http://127.0.0.1:8000/product.html',
                    {params: {q: this.q}}).then(response => {
                    this.myData = response.data
                })
                    .catch(function (error) {
                        console.log(error)
                    })
            },
            // 定义deleteMsg函数
            // 点击删除按钮即可删除当前数据
            // 通过nowIndex确认行数
            deleteMsg: function (n) {
                if (n === -2) {
                    this.myData = []
                } else {
                    this.myData.splice(n, 1)
                }
            }
        }
    }
</script>
