package pages

var IndexHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>拼多多评价生成器</title>
    <link rel="stylesheet" href="https://res.wx.qq.com/open/libs/weui/1.1.3/weui.min.css"/>
    <script src="https://apps.bdimg.com/libs/jquery/2.1.4/jquery.min.js"></script>
</head>
<body>
<style>
    input::-webkit-input-placeholder{
        color:silver;
    }
    input::-moz-placeholder{   /* Mozilla Firefox 19+ */
        color:silver;
    }
    input:-moz-placeholder{    /* Mozilla Firefox 4 to 18 */
        color:silver;
    }
    input:-ms-input-placeholder{  /* Internet Explorer 10-11 */
        color:silver;
    }
</style>
<div class="weui-cells weui-cells_form">
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">商品链接：</label></div>
        <div class="weui-cell__bd">
            <input id="item-link" class="weui-input" type="string" pattern="[0-9]*" placeholder="请输入...">
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">评论字数：</label></div>
        <div class="weui-cell__bd">
            <input id="comment-number" class="weui-input" type="number" pattern="[0-9]*" placeholder="请输入...">
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd"><label class="weui-label">价格折扣：</label></div>
        <div class="weui-cell__bd">
            <input id="comment-discount" class="weui-input" type="number" pattern="[0-9]*" placeholder="例如：0.5">
        </div>
    </div>
    <div class="weui-cells__title">自定义评论头：</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input id="comment-head" class="weui-input" type="text" placeholder="请输入...">
            </div>
        </div>
    </div>
    <div class="weui-cells__title">自定义评论尾：</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input id="comment-foot" class="weui-input" type="text" placeholder="请输入...">
            </div>
        </div>
    </div>
    <div class="weui-cells__title">关键词过滤(使用逗号 , 分隔)：</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input id="comment-filter" class="weui-input" type="text" placeholder="请输入...">
            </div>
        </div>
    </div>
    <div class="weui-cells__title">生成评论：</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <textarea id="comment-result" class="weui-textarea" rows="3" style="font-size: small;" readonly></textarea>
            </div>
        </div>
    </div>
    <div class="weui-cell">
        <div class="weui-cell__hd">
            <canvas width="200" height="150" id="verifyCanvas"></canvas>
            <img id="code_img">
        </div>
        <div class="weui-cell__bd">
            <input id="pic-one-line" class="weui-input" type="string" pattern="[0-9]*" placeholder="请输入第一行">
            <input id="pic-two-line" class="weui-input" type="string" pattern="[0-9]*" placeholder="请输入第二行">
            <input id="pic-three-line" class="weui-input" type="string" pattern="[0-9]*" placeholder="请输入第三行">
            <input id="pic-four-line" class="weui-input" type="string" pattern="[0-9]*" placeholder="默认显示价格">
        </div>
        <div class="weui-cell__ft">
            <a href="javascript:;" class="weui-btn weui-btn_mini weui-btn_primary" onclick="generate()">一键生成</a><br/>
            <a href="javascript:;" id="downloadPic" class="weui-btn weui-btn_mini weui-btn_plain-primary" onclick="savePic()">保存图片</a>
        </div>
    </div>
</div>
<script>
    function generate() {
        let data = {
            item_link: '',
            comment_number: '',
            comment_head: '',
            comment_foot: '',
            comment_filter: '',
            comment_discount: ''
        };
        data.item_link = document.getElementById('item-link').value;
        data.comment_number = document.getElementById('comment-number').value;
        data.comment_head = document.getElementById('comment-head').value;
        data.comment_foot = document.getElementById('comment-foot').value;
        data.comment_filter = document.getElementById('comment-filter').value;
        data.comment_discount = document.getElementById('comment-discount').value;
        if (data.comment_discount.length <= 0) {
            data.comment_discount = '0.5'
        }
        let str = JSON.stringify(data);
        let result = 'generate|||' + str;
        window.external.invoke(result);
    }
    function drawPic(price) {
        let picOne = document.getElementById('pic-one-line').value;
        let picTwo = document.getElementById('pic-two-line').value;
        let picThree = document.getElementById('pic-three-line').value;
        let picFour = document.getElementById('pic-four-line').value;
        if (picFour.length <= 0) {
            picFour = price;
            document.getElementById('pic-four-line').value=price
        }
        drawCode(picOne, picTwo, picThree, picFour);
    }
    function savePic() {
        let canvas = document.getElementById("verifyCanvas");
        let filename = (new Date()).getTime().toString();
        let imgData = canvas.msToBlob();
        let blobObj = new Blob([imgData]);
        window.navigator.msSaveOrOpenBlob(blobObj, filename + '.png')
    }
</script>
<script>
    // 绘制验证码
    function drawCode(picOne, picTwo, picThree, picFour) {
        var canvas = document.getElementById("verifyCanvas"); //获取HTML端画布
        var context = canvas.getContext("2d"); //获取画布2D上下文
        context.fillStyle = "silver"; //画布填充色
        context.fillRect(0, 0, canvas.width, canvas.height); //清空画布
        context.fillStyle = "Black"; //设置字体颜色
        context.font = "30px normal Arial"; //设置字体
        var rand = new Array();
        var x = new Array();
        var y = new Array();
        // for (var i = 0; i < 4; i++) {
        //     rand.push(rand[i]);
        //     rand[i] = nums[Math.floor(Math.random() * nums.length)]
        //     x[i] = i * 20 + 10;
        //     y[i] = Math.random() * 60 + 25;
        //     context.fillText('哈', x[i], y[i]);
        // }
        context.fillText(picFour, 30, 140)
        context.fillText(picThree, 30, 105)
        context.fillText(picTwo, 30, 65)
        context.fillText(picOne, 30, 30)
        //画3条随机线
        for (var i = 0; i < 10; i++) {
            drawline(canvas, context);
        }

        // 画30个随机点
        for (var i = 0; i < 30; i++) {
            drawDot(canvas, context);
        }
        convertCanvasToImage(canvas);
    }

    // 随机线
    function drawline(canvas, context) {
        context.moveTo(Math.floor(Math.random() * canvas.width), Math.floor(Math.random() * canvas.height)); //随机线的起点x坐标是画布x坐标0位置，y坐标是画布高度的随机数
        context.lineTo(Math.floor(Math.random() * canvas.width), Math.floor(Math.random() * canvas.height)); //随机线的终点x坐标是画布宽度，y坐标是画布高度的随机数
        context.lineWidth = 2; //随机线宽
        context.strokeStyle = 'rgba(255,255,255,0.3)'; //随机线描边属性
        context.stroke(); //描边，即起点描到终点
    }

    // 随机点(所谓画点其实就是画1px像素的线，方法不再赘述)
    function drawDot(canvas, context) {
        var px = Math.floor(Math.random() * canvas.width);
        var py = Math.floor(Math.random() * canvas.height);
        context.moveTo(px, py);
        context.lineTo(px + 1, py + 1);
        context.lineWidth = 0.2;
        context.stroke();

    }

    // 绘制图片
    function convertCanvasToImage(canvas) {
        document.getElementById("verifyCanvas").style.display = "none";
        var image = document.getElementById("code_img");
        image.src = canvas.toDataURL("image/png");
        return image;
    }

    // 点击图片刷新
    document.getElementById('code_img').onclick = function () {
        resetCode();
    }

    function resetCode() {
        $('#verifyCanvas').remove();
        $('#code_img').before('<canvas width="200" height="150" id="verifyCanvas"></canvas>')
        drawPic()
    }
</script>
</body>
</html>`
