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
    .weui-cell {
        padding: 2px 5px;
        font-size: 15px;
    }

    .weui-cells__title {
        padding-left: 6px;
        padding-right: 15px;
    }

    input::-webkit-input-placeholder {
        color: silver;
    }

    input::-moz-placeholder { /* Mozilla Firefox 19+ */
        color: silver;
    }

    input:-moz-placeholder { /* Mozilla Firefox 4 to 18 */
        color: silver;
    }

    input:-ms-input-placeholder { /* Internet Explorer 10-11 */
        color: silver;
    }
</style>
<div class="weui-flex">
    <div class="weui-flex__item">
        <div class="weui-cells weui-cells_form">
            <div class="weui-cells__title">生成评论：</div>
            <div class="weui-cells">
                <div class="weui-cell">
                    <div class="weui-cell__bd">
                        <textarea id="comment-result" class="weui-textarea" rows="12" style="font-size: small;" readonly></textarea>
                    </div>
                </div>
            </div>
            <div class="weui-cells__title">生成图片：</div>
            <div class="weui-cell">
                <div class="weui-cell__bd">
                    <canvas width="400" height="330" id="verifyCanvas"></canvas>
                    <img id="code_img">
                </div>
                <div class="weui-cell__ft">
                    <a href="javascript:;" class="weui-btn weui-btn_mini weui-btn_primary" onclick="generate()">一键生成</a><br/>
                    <a href="javascript:;" id="downloadPic" class="weui-btn weui-btn_mini weui-btn_plain-primary"
                       onclick="savePic()">保存图片</a>
                </div>
            </div>
        </div>
    </div>
    <div class="weui-flex__item">
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
            <div class="weui-cells__title">图片设置</div>
            <div class="weui-cells">
                <div class="weui-cell">
                    <div class="weui-cell__hd">
                        <label class="weui-label">修改底图：</label>
                        <div class="weui-uploader__input-box">
                            <input id="uploaderInput" class="weui-uploader__input" type="file" accept="image/*"
                                   multiple="" onchange="changeBaseImage(this.files[0])">
                        </div>
                    </div>
                    <div class="weui-cell__bd">
                        <div class="weui-cell">
                            <div class="weui-cell__hd">
                                手动输入价格：
                            </div>
                            <div class="weui-cell__bd">
                                <input id="pic-price-price" class="weui-input" type="string" pattern="[0-9]*">
                            </div>
                        </div>
                        <div class="weui-cell">
                            <div class="weui-cell__hd">
                                价格X坐标%：
                            </div>
                            <div class="weui-cell__bd">
                                <input id="pic-price-x" class="weui-input" type="string" pattern="[0-9]*">
                            </div>
                        </div>
                        <div class="weui-cell">
                            <div class="weui-cell__hd">
                                价格Y坐标%：
                            </div>
                            <div class="weui-cell__bd">
                                <input id="pic-price-y" class="weui-input" type="string" pattern="[0-9]*">
                            </div>
                        </div>
                        <div class="weui-cell">
                            <div class="weui-cell__hd">
                                价格字体大小：
                            </div>
                            <div class="weui-cell__bd">
                                <input id="pic-price-size" class="weui-input" type="string" pattern="[0-9]*">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    let image = new Image();
    image.src = 'https://img.sucaihuo.com/jquery/32/3287/big.jpg';
    window.onload = function () {
        let canvas = document.getElementById("verifyCanvas");
        let ctx = canvas.getContext('2d');
        ctx.drawImage(image, 0, 0)
    };

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

    function changeBaseImage(file) {
        image.src = window.URL.createObjectURL(file);
        let canvas = document.getElementById("verifyCanvas");
        let ctx = canvas.getContext('2d');
        setTimeout(function () {
            ctx.drawImage(image, 0, 0)
        }, 100);
    }

    function drawPrice(price) {
        let ax = document.getElementById('pic-price-x').value;
        let ay = document.getElementById('pic-price-y').value;
        let s = document.getElementById('pic-price-size').value;
        let p = document.getElementById('pic-price-price').value;
        if (Number(p) > 0 ) {
            price = p
        }
        let canvas = document.getElementById("verifyCanvas");
        let context = canvas.getContext("2d");
        let bx = Number(ax);
        let by = Number(ay);
        if (bx > 100 || bx < 0) {
            bx = 0
        }
        if (by > 100 || by < 0) {
            by = 0
        }
        let x = canvas.width * bx / 100;
        let y = canvas.width * by / 100;
        let size = 'normal bold ' + s + 'px verdana';
        context.fillStyle = "Red";
        context.font = size; //设置字体
        context.drawImage(image, 0, 0);
        context.fillText(price, Number(x), Number(y));
    }

    function savePic() {
        let canvas = document.getElementById("verifyCanvas");
        let filename = (new Date()).getTime().toString();
        let imgData = canvas.msToBlob();
        let blobObj = new Blob([imgData]);
        window.navigator.msSaveOrOpenBlob(blobObj, filename + '.png')
    }
</script>
</body>
</html>`
