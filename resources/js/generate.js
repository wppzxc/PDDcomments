function generate() {
    let data = {
        item_link: '',
        comment_number: '',
        comment_head: '',
        comment_foot: '',
        comment_filter: ''
    };
    data.item_link = document.getElementById('item-link').value;
    data.comment_number = document.getElementById('comment-number').value;
    data.comment_head = document.getElementById('comment-head').value;
    data.comment_foot = document.getElementById('comment-foot').value;
    data.comment_filter = document.getElementById('comment-filter').value;
    let str = JSON.stringify(data);
    let result = 'generate|||' + str;
    window.external.invoke(result);
    console.log(result)
}

!function () {
    document.getElementById("comment-result").value = '%s'
}();

!function () {
    drawPic('%s')
}();