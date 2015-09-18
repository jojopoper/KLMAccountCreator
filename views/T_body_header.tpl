{{define "T_body_header"}}
<header class="header dark-bg">
    <div class="toggle-nav">
        <div class="icon-reorder tooltips" data-original-title="Toggle Navigation" data-placement="bottom"></div>
    </div>

    <!--logo start-->
    <a href="http://klm.strllar.org/" class="logo">Strllar <span class="lite">KLM</span></a>
    <!--logo end-->

    <div class="top-nav notification-row pull-right">
        <div class="btn-row">
            <div class="btn-group">
	            <form method="POST">
	              <button type="submit" name="languageBtn" value="CH" class="{{if .L_IsChinese}}active{{end}} btn btn-primary">中文</button>
	              <button type="submit" name="languageBtn" value="EN" class="{{if .L_IsEnglish}}active{{end}} btn btn-primary">English</button>
              </form>
            </div>
        </div>
    </div>
</header>
{{end}}