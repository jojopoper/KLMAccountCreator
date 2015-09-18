{{define "T_body_sidebar"}}
<aside>
    <div id="sidebar"  class="nav-collapse ">
        <!-- sidebar menu start-->
        <ul class="sidebar-menu">
            <li class="">
                <a class="" href="http://klm.strllar.org/">
                    <i class="icon_house_alt"></i>
                    <span>
                    {{if .L_IsChinese}}
                    主页
                    {{else}}
                    Home
                    {{end}}
                    </span>
                </a>
            </li>
            <li class="sub-menu">
                <a href="javascript:;" class="">
                    <i class="icon_document_alt"></i>
                    <span>
                    {{if .L_IsChinese}}
                    账户
                    {{else}}
                    Account
                    {{end}}
                    </span>
                    <span class="menu-arrow arrow_carrot-right"></span>
                </a>
                <ul class="sub">
                    <li><a class="" href="#">
                    {{if .L_IsChinese}}
                    新建
                    {{else}}
                    New
                    {{end}}
                    </a></li>
                </ul>
            </li>
        </ul>
        <!-- sidebar menu end-->
    </div>
</aside>
{{end}}