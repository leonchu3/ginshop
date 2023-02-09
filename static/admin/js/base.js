$(function(){
    baseApp.init();
})

var baseApp={
    init:function(){
        this.initAside()
		this.confirmDelete()
    },
    //侧边栏滑动收缩
    initAside:function(){
        $('.aside h4').click(function(){
			$(this).siblings('ul').slideToggle();
		}) 
    },
    //设置iframe的高度 
	resizeIframe:function(){					
		$("#rightMain").height($(window).height()-80)
	},
    confirmDelete:function(){
        $(".delete").click(function(){
			var flag=confirm("您确定要删除了吗？")
			return flag
        })
    }
}