package main

import (
	"fmt"
	"log"
	"net/http"
)

const htmlContent = `<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>JSON Toolbox</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: #1e1e2e;
            color: #cdd6f4;
            padding: 20px 30px;
            min-height: 100vh;
        }
        h1 { color: #89b4fa; text-align: center; margin-bottom: 6px; font-size: 24px; }
        .subtitle { text-align: center; color: #6c7086; margin-bottom: 20px; font-size: 13px; }

        .main-tabs {
            display: flex;
            justify-content: center;
            gap: 4px;
            margin-bottom: 20px;
            flex-wrap: wrap;
        }
        .main-tab {
            padding: 12px 18px;
            background: #313244;
            color: #a6adc8;
            border: 1px solid #45475a;
            border-radius: 10px 10px 0 0;
            cursor: pointer;
            font-size: 13px;
            font-weight: 700;
            transition: all 0.15s;
            border-bottom: none;
        }
        .main-tab:hover { color: #cdd6f4; background: #45475a; }
        .main-tab.active { background: #89b4fa; color: #1e1e2e; border-color: #89b4fa; }

        .tool-panel { display: none; }
        .tool-panel.active { display: block; }

        .workspace {
            max-width: 1500px;
            margin: 0 auto;
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
        }
        .panel {
            background: #313244;
            padding: 18px;
            border-radius: 12px;
            box-shadow: 0 4px 15px rgba(0,0,0,0.3);
            display: flex;
            flex-direction: column;
        }
        .panel-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 10px;
            padding-bottom: 10px;
            border-bottom: 1px solid #45475a;
        }
        .panel-title { color: #89b4fa; font-weight: 600; font-size: 13px; text-transform: uppercase; letter-spacing: 0.5px; }
        .stats { color: #6c7086; font-size: 12px; }

        textarea {
            flex: 1;
            width: 100%;
            min-height: 450px;
            background: #181825;
            color: #cdd6f4;
            border: 1px solid #45475a;
            border-radius: 8px;
            padding: 14px;
            font-family: 'Consolas', 'Monaco', monospace;
            font-size: 13px;
            resize: none;
            line-height: 1.5;
        }
        textarea:focus { outline: 2px solid #89b4fa; border-color: transparent; }
        textarea[readonly] { color: #a6e3a1; }

        .toolbar { display: flex; gap: 8px; margin-top: 10px; flex-wrap: wrap; align-items: center; }
        button {
            padding: 8px 14px;
            border: none;
            border-radius: 6px;
            font-size: 13px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.15s;
            display: flex;
            align-items: center;
            gap: 5px;
            white-space: nowrap;
        }
        .btn-primary { background: #89b4fa; color: #1e1e2e; }
        .btn-primary:hover { background: #b4befe; }
        .btn-success { background: #a6e3a1; color: #1e1e2e; }
        .btn-success:hover { background: #94e2d5; }
        .btn-warning { background: #fab387; color: #1e1e2e; }
        .btn-warning:hover { background: #f9e2af; }
        .btn-danger { background: #f38ba8; color: #1e1e2e; }
        .btn-danger:hover { background: #eba0ac; }
        .btn-secondary { background: #45475a; color: #cdd6f4; }
        .btn-secondary:hover { background: #585b70; }
        .btn-swap { background: #cba6f7; color: #1e1e2e; }
        .btn-swap:hover { background: #b4befe; }

        input[type="text"] {
            padding: 8px 12px;
            background: #181825;
            color: #cdd6f4;
            border: 1px solid #45475a;
            border-radius: 6px;
            font-size: 13px;
            width: 140px;
        }
        input[type="text"]:focus { outline: 2px solid #89b4fa; border-color: transparent; }

        .sub-tabs { display: flex; gap: 2px; margin-bottom: 10px; }
        .sub-tab {
            padding: 6px 14px;
            background: #181825;
            color: #a6adc8;
            border: 1px solid #45475a;
            border-radius: 6px 6px 0 0;
            cursor: pointer;
            font-size: 12px;
            font-weight: 600;
            transition: all 0.15s;
        }
        .sub-tab.active { background: #313244; color: #89b4fa; border-color: #89b4fa; }
        .sub-tab:hover:not(.active) { color: #cdd6f4; }
        .sub-content { display: none; flex: 1; }
        .sub-content.active { display: block; }

        .tree-view {
            background: #181825;
            border: 1px solid #45475a;
            border-radius: 8px;
            padding: 14px;
            overflow: auto;
            max-height: 450px;
            font-family: 'Consolas', monospace;
            font-size: 13px;
        }
        .tree-node { margin-left: 20px; line-height: 1.6; }
        .tree-node.root { margin-left: 0; }
        .tree-key { color: #89b4fa; font-weight: 600; }
        .tree-string { color: #a6e3a1; cursor: pointer; }
        .tree-string:hover { background: #45475a; border-radius: 3px; }
        .tree-number { color: #fab387; cursor: pointer; }
        .tree-number:hover { background: #45475a; border-radius: 3px; }
        .tree-boolean { color: #cba6f7; font-weight: 600; }
        .tree-null { color: #f38ba8; font-style: italic; }
        .tree-bracket { color: #6c7086; }
        .tree-toggle { cursor: pointer; user-select: none; display: inline-block; width: 14px; color: #6c7086; font-size: 11px; }
        .tree-toggle:hover { color: #89b4fa; }
        .tree-children { margin-left: 8px; }
        .tree-children.collapsed { display: none; }
        .tree-summary { color: #6c7086; font-style: italic; font-size: 12px; }
        .raw-view {
            background: #181825;
            color: #cdd6f4;
            border: 1px solid #45475a;
            border-radius: 8px;
            padding: 14px;
            font-family: 'Consolas', monospace;
            font-size: 13px;
            white-space: pre-wrap;
            overflow: auto;
            max-height: 450px;
            line-height: 1.5;
        }

        .changes-log {
            background: #181825;
            border: 1px solid #45475a;
            border-radius: 8px;
            padding: 12px;
            margin-top: 10px;
            max-height: 180px;
            overflow-y: auto;
            font-family: 'Consolas', monospace;
            font-size: 12px;
            display: none;
        }
        .changes-log.active { display: block; }
        .change-item { padding: 3px 0; border-bottom: 1px solid #313244; display: flex; gap: 8px; align-items: center; }
        .change-item:last-child { border-bottom: none; }
        .old-key { color: #f38ba8; text-decoration: line-through; }
        .arrow { color: #6c7086; }
        .new-key { color: #a6e3a1; }
        .field-id { color: #6c7086; font-size: 11px; margin-left: auto; }

        .status {
            position: fixed; top: 20px; right: 20px;
            padding: 12px 20px; border-radius: 8px;
            font-weight: 600; font-size: 14px;
            box-shadow: 0 4px 15px rgba(0,0,0,0.3);
            opacity: 0; transform: translateX(100px);
            transition: all 0.3s; z-index: 1000;
        }
        .status.show { opacity: 1; transform: translateX(0); }
        .status.success { background: #a6e3a1; color: #1e1e2e; }
        .status.error { background: #f38ba8; color: #1e1e2e; }

        .info-box {
            background: #181825;
            border-left: 3px solid #89b4fa;
            padding: 10px 14px;
            border-radius: 6px;
            margin-bottom: 16px;
            font-size: 13px;
            color: #a6adc8;
            max-width: 1500px;
            margin: 0 auto 16px auto;
        }
        .info-box code { background: #313244; padding: 1px 6px; border-radius: 3px; color: #f38ba8; font-family: 'Consolas', monospace; }
        .info-box strong { color: #89b4fa; }

        .mode-switch {
            display: inline-flex;
            background: #181825;
            border: 1px solid #45475a;
            border-radius: 8px;
            padding: 3px;
            gap: 2px;
        }
        .mode-btn {
            padding: 6px 16px;
            background: transparent;
            color: #a6adc8;
            border: none;
            border-radius: 6px;
            font-size: 13px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.15s;
        }
        .mode-btn.active { background: #89b4fa; color: #1e1e2e; }
        .mode-btn:hover:not(.active) { color: #cdd6f4; }

        /* Cookie specific */
        .cookie-result {
            background: #181825;
            border: 1px solid #45475a;
            border-radius: 8px;
            padding: 14px;
            font-family: 'Consolas', monospace;
            font-size: 13px;
            color: #a6e3a1;
            word-break: break-all;
            min-height: 100px;
            max-height: 450px;
            overflow: auto;
            line-height: 1.5;
        }

        @media (max-width: 1100px) { .workspace { grid-template-columns: 1fr; } }
    </style>
</head>
<body>
    <h1>🧰 JSON Toolbox</h1>
    <p class="subtitle">6 инструментов в одном месте — всё локально, ничего не уходит в интернет</p>

    <div class="main-tabs">
        <div class="main-tab active" onclick="switchMainTab('formatter', event)">✨ Formatter</div>
        <div class="main-tab" onclick="switchMainTab('keyformatter', event)">🔑 Key Formatter</div>
        <div class="main-tab" onclick="switchMainTab('base64', event)">🔐 Base64</div>
        <div class="main-tab" onclick="switchMainTab('url', event)">🌐 URL</div>
        <div class="main-tab" onclick="switchMainTab('idcleaner', event)">🧹 ID Cleaner</div>
        <div class="main-tab" onclick="switchMainTab('cookie', event)">🍪 Cookie</div>
    </div>

    <!-- ====== TAB 1: JSON FORMATTER ====== -->
    <div id="formatterPanel" class="tool-panel active">
        <div class="workspace">
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📥 Исходный JSON</span>
                    <span class="stats" id="fInputStats">0 символов</span>
                </div>
                <textarea id="fInput" placeholder='Вставь сюда JSON...'></textarea>
                <div class="toolbar">
                    <button class="btn-primary" onclick="fFormat()">✨ Форматировать</button>
                    <button class="btn-warning" onclick="fMinify()">📦 Минифицировать</button>
                    <button class="btn-danger" onclick="fClear()">🗑️ Очистить</button>
                </div>
            </div>
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📤 Результат</span>
                    <span class="stats" id="fOutputStats">—</span>
                </div>
                <div class="sub-tabs">
                    <div class="sub-tab active" onclick="fSwitchSub('tree',this)">🌳 Tree View</div>
                    <div class="sub-tab" onclick="fSwitchSub('raw',this)">📄 Raw JSON</div>
                </div>
                <div id="fTreeSub" class="sub-content active">
                    <div class="toolbar" style="margin-top:0;margin-bottom:8px;">
                        <button class="btn-secondary" onclick="fExpandAll()">➕ Развернуть</button>
                        <button class="btn-secondary" onclick="fCollapseAll()">➖ Свернуть</button>
                    </div>
                    <div class="tree-view" id="fTreeView"><span style="color:#6c7086;font-style:italic;">Нажми "Форматировать"</span></div>
                </div>
                <div id="fRawSub" class="sub-content">
                    <div class="raw-view" id="fRawView"><span style="color:#6c7086;font-style:italic;">Нажми "Форматировать"</span></div>
                </div>
                <div class="toolbar" style="margin-top:12px;">
                    <input type="text" id="fFilename" value="data" placeholder="Имя файла">
                    <span style="color:#a6adc8;">.json</span>
                    <button class="btn-success" onclick="fDownload()">⬇️ Скачать</button>
                    <button class="btn-secondary" onclick="fCopy()">📋 Копировать</button>
                </div>
            </div>
        </div>
    </div>

    <!-- ====== TAB 2: KEY FORMATTER ====== -->
    <div id="keyformatterPanel" class="tool-panel">
        <div class="info-box">
            <strong>📋 Правило:</strong> Рекурсивно обходит <strong>весь JSON</strong>. Находит все объекты с <code>data.key</code> (строка) и заменяет <code>__</code> на <code>.</code>. Структура сохраняется.
        </div>
        <div class="workspace">
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📥 Исходный JSON</span>
                    <span class="stats" id="kInputStats">0 символов</span>
                </div>
                <textarea id="kInput" placeholder='Вставь сюда JSON с data.key...'></textarea>
                <div class="toolbar">
                    <button class="btn-primary" onclick="kProcess()">⚙️ Обработать</button>
                    <button class="btn-danger" onclick="kClear()">🗑️ Очистить</button>
                </div>
            </div>
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📤 Результат</span>
                    <span class="stats" id="kOutputStats">—</span>
                </div>
                <textarea id="kOutput" readonly placeholder='Обработанный JSON...'></textarea>
                <div id="kChangesLog" class="changes-log"></div>
                <div class="toolbar">
                    <input type="text" id="kFilename" value="processed" placeholder="Имя файла">
                    <span style="color:#a6adc8;">.json</span>
                    <button class="btn-success" onclick="kDownload()">⬇️ Скачать</button>
                    <button class="btn-secondary" onclick="kCopy()">📋 Копировать</button>
                </div>
            </div>
        </div>
    </div>

    <!-- ====== TAB 3: BASE64 ====== -->
    <div id="base64Panel" class="tool-panel">
        <div class="info-box">
            <strong>📋 Base64</strong> — кодирование/декодирование. Полная поддержка UTF-8 (кириллица, эмодзи).
        </div>
        <div class="workspace">
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📥 Исходный текст</span>
                    <span class="stats" id="bInputStats">0 символов</span>
                </div>
                <textarea id="bInput" placeholder='Текст для кодирования или Base64 для декодирования...'></textarea>
                <div class="toolbar">
                    <div class="mode-switch">
                        <button class="mode-btn active" id="bModeEnc" onclick="bSetMode('enc')">🔒 Encode</button>
                        <button class="mode-btn" id="bModeDec" onclick="bSetMode('dec')">🔓 Decode</button>
                    </div>
                    <button class="btn-primary" onclick="bProcess()">⚙️ Обработать</button>
                    <button class="btn-swap" onclick="bSwap()">🔄 Swap</button>
                    <button class="btn-danger" onclick="bClear()">🗑️ Очистить</button>
                </div>
            </div>
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📤 Результат</span>
                    <span class="stats" id="bOutputStats">—</span>
                </div>
                <textarea id="bOutput" readonly placeholder='Результат...'></textarea>
                <div class="toolbar">
                    <button class="btn-success" onclick="bCopy()">📋 Копировать</button>
                    <button class="btn-secondary" onclick="bDownload()">⬇️ Скачать</button>
                </div>
            </div>
        </div>
    </div>

    <!-- ====== TAB 4: URL ENCODER ====== -->
    <div id="urlPanel" class="tool-panel">
        <div class="info-box">
            <strong>📋 URL Encoder</strong> — кодирует/декодирует строку для безопасного использования в URL.
        </div>
        <div class="workspace">
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📥 Исходный текст</span>
                    <span class="stats" id="uInputStats">0 символов</span>
                </div>
                <textarea id="uInput" placeholder='URL или текст...'></textarea>
                <div class="toolbar">
                    <div class="mode-switch">
                        <button class="mode-btn active" id="uModeEnc" onclick="uSetMode('enc')">🔒 Encode</button>
                        <button class="mode-btn" id="uModeDec" onclick="uSetMode('dec')">🔓 Decode</button>
                    </div>
                    <button class="btn-primary" onclick="uProcess()">⚙️ Обработать</button>
                    <button class="btn-swap" onclick="uSwap()">🔄 Swap</button>
                    <button class="btn-danger" onclick="uClear()">🗑️ Очистить</button>
                </div>
            </div>
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📤 Результат</span>
                    <span class="stats" id="uOutputStats">—</span>
                </div>
                <textarea id="uOutput" readonly placeholder='Результат...'></textarea>
                <div class="toolbar">
                    <button class="btn-success" onclick="uCopy()">📋 Копировать</button>
                    <button class="btn-secondary" onclick="uDownload()">⬇️ Скачать</button>
                </div>
            </div>
        </div>
    </div>

    <!-- ====== TAB 5: ID CLEANER ====== -->
    <div id="idcleanerPanel" class="tool-panel">
        <div class="info-box">
            <strong>📋 Правило:</strong> Берёт массив <code>requests[]</code>, рекурсивно удаляет все <code>id</code> и <code>versionId</code>, оборачивает результат в <code>{ "testRequests": [...] }</code>.
        </div>
        <div class="workspace">
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📥 Исходный JSON</span>
                    <span class="stats" id="tcInputStats">0 символов</span>
                </div>
                <textarea id="tcInput" placeholder='Вставь сюда JSON с requests[]...'></textarea>
                <div class="toolbar">
                    <button class="btn-primary" onclick="tcProcess()">🧹 Обработать</button>
                    <button class="btn-danger" onclick="tcClear()">🗑️ Очистить</button>
                </div>
            </div>
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📤 Результат (testRequests)</span>
                    <span class="stats" id="tcOutputStats">—</span>
                </div>
                <textarea id="tcOutput" readonly placeholder='Здесь появится JSON с testRequests[]...'></textarea>
                <div id="tcLog" class="changes-log"></div>
                <div class="toolbar">
                    <input type="text" id="tcFilename" value="test_requests" placeholder="Имя файла">
                    <span style="color:#a6adc8;">.json</span>
                    <button class="btn-success" onclick="tcDownload()">⬇️ Скачать</button>
                    <button class="btn-secondary" onclick="tcCopy()">📋 Копировать</button>
                </div>
            </div>
        </div>
    </div>

    <!-- ====== TAB 6: COOKIE EXTRACTOR ====== -->
    <div id="cookiePanel" class="tool-panel">
        <div class="info-box">
            <strong>📋 Cookie Extractor:</strong> Вставь строку cookie из браузера, укажи имя нужной куки — получишь значение. По умолчанию ищет <code>authToken_production</code>.
        </div>
        <div class="workspace">
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📥 Строка Cookie</span>
                    <span class="stats" id="ceInputStats">0 символов</span>
                </div>
                <textarea id="ceInput" placeholder='Вставь сюда строку cookie из браузера... Например: tmr_lvid=abc; authToken_production=eyJhbGc...'></textarea>
                <div class="toolbar">
                    <button class="btn-primary" onclick="ceExtract()">🍪 Извлечь</button>
                    <button class="btn-danger" onclick="ceClear()">🗑️ Очистить</button>
                </div>
            </div>
            <div class="panel">
                <div class="panel-header">
                    <span class="panel-title">📤 Результат</span>
                    <span class="stats" id="ceOutputStats">—</span>
                </div>
                <div class="toolbar" style="margin-top:0;margin-bottom:10px;">
                    <label style="color:#a6adc8;font-weight:600;">Имя cookie:</label>
                    <input type="text" id="ceCookieName" value="authToken_production" style="width:220px;">
                </div>
                <div class="cookie-result" id="ceOutput">
                    <span style="color:#6c7086;font-style:italic;">Нажми "Извлечь", чтобы получить значение</span>
                </div>
                <div class="toolbar" style="margin-top:12px;">
                    <button class="btn-success" onclick="ceCopy()">📋 Копировать</button>
                </div>
            </div>
        </div>
    </div>

    <div id="status" class="status"></div>

    <script>
    /* === GLOBAL === */
    function showStatus(msg, type) {
        const s = document.getElementById('status');
        s.textContent = msg; s.className = 'status show ' + type;
        setTimeout(() => s.classList.remove('show'), 2500);
    }
    function esc(s) { return String(s).replace(/&/g,'&amp;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;'); }
    function dlText(text, name, mime) {
        let fn = name.trim() || 'data';
        const a = document.createElement('a');
        a.href = URL.createObjectURL(new Blob([text], {type: mime||'text/plain'}));
        a.download = fn; document.body.appendChild(a); a.click();
        document.body.removeChild(a);
        showStatus('⬇️ "' + fn + '" скачан!', 'success');
    }
    function cp(text) { navigator.clipboard.writeText(text).then(() => showStatus('📋 Скопировано!','success')); }

    function switchMainTab(name, ev) {
        document.querySelectorAll('.main-tab').forEach(t => t.classList.remove('active'));
        document.querySelectorAll('.tool-panel').forEach(p => p.classList.remove('active'));
        ev.target.classList.add('active');
        document.getElementById(name + 'Panel').classList.add('active');
    }

    /* === TAB 1: JSON FORMATTER === */
    let fParsed=null, fStr='';
    document.getElementById('fInput').addEventListener('input', function(){ document.getElementById('fInputStats').textContent=this.value.length+' символов'; });
    function fFormat(){
        const v=document.getElementById('fInput').value.trim();
        if(!v){showStatus('⚠️ Введи JSON!','error');return;}
        try{ fParsed=JSON.parse(v); fStr=JSON.stringify(fParsed,null,4);
            document.getElementById('fRawView').textContent=fStr; fRenderTree(fParsed);
            document.getElementById('fOutputStats').textContent=fStr.length+' символов • '+fCount(fParsed)+' элементов';
            showStatus('✅ Отформатировано!','success');
        }catch(e){showStatus('❌ '+e.message,'error');}
    }
    function fMinify(){
        const v=document.getElementById('fInput').value.trim(); if(!v)return;
        try{ fParsed=JSON.parse(v); fStr=JSON.stringify(fParsed);
            document.getElementById('fRawView').textContent=fStr;
            document.getElementById('fInput').value=fStr; fRenderTree(fParsed);
            document.getElementById('fOutputStats').textContent=fStr.length+' символов';
            showStatus('📦 Минифицировано!','success');
        }catch(e){showStatus('❌ '+e.message,'error');}
    }
    function fClear(){
        document.getElementById('fInput').value='';
        document.getElementById('fRawView').innerHTML='<span style="color:#6c7086;font-style:italic;">Нажми "Форматировать"</span>';
        document.getElementById('fTreeView').innerHTML='<span style="color:#6c7086;font-style:italic;">Нажми "Форматировать"</span>';
        document.getElementById('fInputStats').textContent='0 символов';
        document.getElementById('fOutputStats').textContent='—';
        fParsed=null;fStr=''; showStatus('🗑️ Очищено!','success');
    }
    function fSwitchSub(name,el){
        el.parentElement.querySelectorAll('.sub-tab').forEach(t=>t.classList.remove('active'));
        el.classList.add('active');
        document.getElementById('fTreeSub').classList.toggle('active',name==='tree');
        document.getElementById('fRawSub').classList.toggle('active',name==='raw');
    }
    function fDownload(){if(!fStr){showStatus('⚠️ Сначала отформатируй!','error');return;} dlText(fStr,(document.getElementById('fFilename').value||'data')+'.json','application/json');}
    function fCopy(){if(!fStr){showStatus('⚠️ Нечего копировать!','error');return;} cp(fStr);}
    function fCount(o){let c=0;(function t(v){if(Array.isArray(v)){c+=v.length;v.forEach(t);}else if(v&&typeof v==='object'){c+=Object.keys(v).length;Object.values(v).forEach(t);}})(o);return c;}
    function fExpandAll(){document.querySelectorAll('#fTreeView .tree-children').forEach(e=>e.classList.remove('collapsed'));document.querySelectorAll('#fTreeView .tree-toggle').forEach(e=>e.textContent='▼');document.querySelectorAll('#fTreeView .tree-summary').forEach(e=>e.style.display='none');}
    function fCollapseAll(){document.querySelectorAll('#fTreeView .tree-children').forEach(e=>e.classList.add('collapsed'));document.querySelectorAll('#fTreeView .tree-toggle').forEach(e=>e.textContent='▶');document.querySelectorAll('#fTreeView .tree-summary').forEach(e=>e.style.display='inline');}
    function fRenderTree(data){const c=document.getElementById('fTreeView');c.innerHTML='';c.appendChild(fBuildNode(null,data,true));}
    function fBuildNode(key,val,root){
        const n=document.createElement('div'); n.className='tree-node'+(root?' root':'');
        if(val===null){n.innerHTML=(key!==null?'<span class="tree-key">"'+esc(key)+'"</span>: ':'')+'<span class="tree-null">null</span>';return n;}
        const t=Array.isArray(val)?'array':typeof val;
        if(t==='array'||t==='object') return fBuildCollapsible(n,key,val,t==='array'?'[':'{',t==='array'?']':'}',t==='array');
        let d='',cls='';
        if(t==='string'){d='"'+esc(val)+'"';cls='tree-string';}
        else if(t==='number'){d=val;cls='tree-number';}
        else if(t==='boolean'){d=val;cls='tree-boolean';}
        n.innerHTML=(key!==null?'<span class="tree-key">"'+esc(key)+'"</span>: ':'')+'<span class="'+cls+'">'+d+'</span>';
        return n;
    }
    function fBuildCollapsible(node,key,val,ob,cb,isArr){
        const entries=isArr?val:Object.entries(val);
        const cnt=isArr?val.length:Object.keys(val).length;
        const hdr=document.createElement('span');hdr.style.cursor='pointer';hdr.style.userSelect='none';
        const tog=document.createElement('span');tog.className='tree-toggle';tog.textContent='▼';
        const sum=document.createElement('span');sum.className='tree-summary';sum.style.display='none';
        sum.textContent=' '+cnt+(isArr?' items':(cnt===1?' key':' keys'));
        hdr.innerHTML=(key!==null?'<span class="tree-key">"'+esc(key)+'"</span>: ':'')+'<span class="tree-bracket">'+ob+'</span>';
        hdr.insertBefore(tog,hdr.firstChild);hdr.appendChild(sum);
        node.appendChild(hdr);
        const ch=document.createElement('div');ch.className='tree-children';
        if(isArr) val.forEach((v,i)=>ch.appendChild(fBuildNode(i,v)));
        else entries.forEach(([k,v])=>ch.appendChild(fBuildNode(k,v)));
        const cl=document.createElement('div');cl.innerHTML='<span class="tree-bracket">'+cb+'</span>';cl.style.marginLeft='20px';
        node.appendChild(ch);node.appendChild(cl);
        hdr.addEventListener('click',e=>{e.stopPropagation();const c=ch.classList.toggle('collapsed');tog.textContent=c?'▶':'▼';sum.style.display=c?'inline':'none';});
        return node;
    }

    /* === TAB 2: KEY FORMATTER === */
    let kStr='';
    document.getElementById('kInput').addEventListener('input',function(){document.getElementById('kInputStats').textContent=this.value.length+' символов';});
    function kProcess(){
        const v=document.getElementById('kInput').value.trim();
        if(!v){showStatus('⚠️ Введи JSON!','error');return;}
        let data; try{data=JSON.parse(v);}catch(e){showStatus('❌ '+e.message,'error');return;}
        const changes=[]; let count=0;
        function walk(obj){
            if(Array.isArray(obj)){obj.forEach(item=>walk(item));}
            else if(obj!==null&&typeof obj==='object'){
                if(obj.data&&typeof obj.data==='object'&&!Array.isArray(obj.data)&&typeof obj.data.key==='string'){
                    const old=obj.data.key, nw=old.replace(/__/g,'.');
                    if(old!==nw){obj.data.key=nw;changes.push({id:obj.id||'—',old:old,new:nw});count++;}
                }
                Object.values(obj).forEach(val=>walk(val));
            }
        }
        walk(data);
        kStr=JSON.stringify(data,null,4);
        document.getElementById('kOutput').value=kStr;
        document.getElementById('kOutputStats').textContent=kStr.length+' символов • изменено: '+count;
        const log=document.getElementById('kChangesLog');
        if(changes.length===0){log.classList.remove('active');}
        else{
            let html='<div style="color:#89b4fa;font-weight:600;margin-bottom:6px;">📝 Изменения ('+changes.length+'):</div>';
            changes.forEach(c=>{html+='<div class="change-item"><span class="old-key">'+esc(c.old)+'</span><span class="arrow">→</span><span class="new-key">'+esc(c.new)+'</span><span class="field-id">#'+c.id+'</span></div>';});
            log.innerHTML=html;log.classList.add('active');
        }
        showStatus(count>0?'✅ Обработано ключей: '+count:'ℹ️ Ключей с "__" не найдено','success');
    }
    function kClear(){
        document.getElementById('kInput').value='';document.getElementById('kOutput').value='';
        document.getElementById('kInputStats').textContent='0 символов';
        document.getElementById('kOutputStats').textContent='—';
        document.getElementById('kChangesLog').classList.remove('active');
        kStr='';showStatus('🗑️ Очищено!','success');
    }
    function kDownload(){if(!kStr){showStatus('⚠️ Сначала обработай!','error');return;} dlText(kStr,(document.getElementById('kFilename').value||'processed')+'.json','application/json');}
    function kCopy(){if(!kStr){showStatus('⚠️ Нечего копировать!','error');return;} cp(kStr);}

    /* === TAB 3: BASE64 === */
    let bMode='enc';
    document.getElementById('bInput').addEventListener('input',function(){document.getElementById('bInputStats').textContent=this.value.length+' символов';});
    function bSetMode(m){bMode=m;document.getElementById('bModeEnc').classList.toggle('active',m==='enc');document.getElementById('bModeDec').classList.toggle('active',m==='dec');}
    function bProcess(){
        const input=document.getElementById('bInput').value;
        if(!input){showStatus('⚠️ Введи текст!','error');return;}
        try{
            let result;
            if(bMode==='enc'){const bytes=new TextEncoder().encode(input);let bin='';bytes.forEach(b=>bin+=String.fromCharCode(b));result=btoa(bin);}
            else{const bin=atob(input);const bytes=new Uint8Array(bin.length);for(let i=0;i<bin.length;i++)bytes[i]=bin.charCodeAt(i);result=new TextDecoder().decode(bytes);}
            document.getElementById('bOutput').value=result;
            document.getElementById('bOutputStats').textContent=result.length+' символов';
            showStatus('✅ Base64 '+(bMode==='enc'?'закодирован':'декодирован')+'!','success');
        }catch(e){showStatus('❌ Ошибка: '+e.message,'error');}
    }
    function bSwap(){const a=document.getElementById('bInput'),b=document.getElementById('bOutput');const t=a.value;a.value=b.value;b.value=t;bSetMode(bMode==='enc'?'dec':'enc');showStatus('🔄 Swap!','success');}
    function bClear(){document.getElementById('bInput').value='';document.getElementById('bOutput').value='';document.getElementById('bInputStats').textContent='0 символов';document.getElementById('bOutputStats').textContent='—';showStatus('🗑️ Очищено!','success');}
    function bCopy(){const v=document.getElementById('bOutput').value;if(!v){showStatus('⚠️ Нечего копировать!','error');return;} cp(v);}
    function bDownload(){const v=document.getElementById('bOutput').value;if(!v){showStatus('⚠️ Нечего сохранять!','error');return;} dlText(v,'base64_'+(bMode==='enc'?'encoded':'decoded')+'.txt','text/plain');}

    /* === TAB 4: URL === */
    let uMode='enc';
    document.getElementById('uInput').addEventListener('input',function(){document.getElementById('uInputStats').textContent=this.value.length+' символов';});
    function uSetMode(m){uMode=m;document.getElementById('uModeEnc').classList.toggle('active',m==='enc');document.getElementById('uModeDec').classList.toggle('active',m==='dec');}
    function uProcess(){
        const input=document.getElementById('uInput').value;
        if(!input){showStatus('⚠️ Введи текст!','error');return;}
        try{
            let result = uMode==='enc' ? encodeURIComponent(input) : decodeURIComponent(input);
            document.getElementById('uOutput').value=result;
            document.getElementById('uOutputStats').textContent=result.length+' символов';
            showStatus('✅ URL '+(uMode==='enc'?'закодирован':'декодирован')+'!','success');
        }catch(e){showStatus('❌ Ошибка: '+e.message,'error');}
    }
    function uSwap(){const a=document.getElementById('uInput'),b=document.getElementById('uOutput');const t=a.value;a.value=b.value;b.value=t;uSetMode(uMode==='enc'?'dec':'enc');showStatus('🔄 Swap!','success');}
    function uClear(){document.getElementById('uInput').value='';document.getElementById('uOutput').value='';document.getElementById('uInputStats').textContent='0 символов';document.getElementById('uOutputStats').textContent='—';showStatus('🗑️ Очищено!','success');}
    function uCopy(){const v=document.getElementById('uOutput').value;if(!v){showStatus('⚠️ Нечего копировать!','error');return;} cp(v);}
    function uDownload(){const v=document.getElementById('uOutput').value;if(!v){showStatus('⚠️ Нечего сохранять!','error');return;} dlText(v,'url_'+(uMode==='enc'?'encoded':'decoded')+'.txt','text/plain');}

    /* === TAB 5: ID CLEANER === */
    let tcStr='';
    document.getElementById('tcInput').addEventListener('input',function(){document.getElementById('tcInputStats').textContent=this.value.length+' символов';});
    function tcProcess(){
        const v=document.getElementById('tcInput').value.trim();
        if(!v){showStatus('⚠️ Введи JSON!','error');return;}
        let data;
        try{data=JSON.parse(v);}catch(e){showStatus('❌ Невалидный JSON: '+e.message,'error');return;}
        let requests;
        if(Array.isArray(data)){requests=data;}
        else if(data.requests&&Array.isArray(data.requests)){requests=data.requests;}
        else if(data.testRequests&&Array.isArray(data.testRequests)){requests=data.testRequests;}
        else{showStatus('❌ Не найден массив requests[] или testRequests[]','error');return;}
        let removedCount=0;
        function removeIds(obj){
            if(Array.isArray(obj)){obj.forEach(item=>removeIds(item));}
            else if(obj!==null&&typeof obj==='object'){
                if('id' in obj){delete obj.id;removedCount++;}
                if('versionId' in obj){delete obj.versionId;removedCount++;}
                Object.values(obj).forEach(val=>removeIds(val));
            }
        }
        const cleanRequests=JSON.parse(JSON.stringify(requests));
        removeIds(cleanRequests);
        const result={testRequests:cleanRequests};
        tcStr=JSON.stringify(result,null,4);
        document.getElementById('tcOutput').value=tcStr;
        document.getElementById('tcOutputStats').textContent=tcStr.length+' символов • удалено полей: '+removedCount;
        const log=document.getElementById('tcLog');
        let html='<div style="color:#89b4fa;font-weight:600;margin-bottom:6px;">🧹 Результат обработки:</div>';
        html+='<div class="change-item"><span class="new-key">✅ requests[] → testRequests[]</span></div>';
        html+='<div class="change-item"><span class="new-key">🗑️ Удалено полей id/versionId: '+removedCount+'</span></div>';
        html+='<div class="change-item"><span class="new-key">📦 Объектов в testRequests: '+cleanRequests.length+'</span></div>';
        log.innerHTML=html;log.classList.add('active');
        showStatus('✅ Готово! Удалено полей: '+removedCount,'success');
    }
    function tcClear(){
        document.getElementById('tcInput').value='';document.getElementById('tcOutput').value='';
        document.getElementById('tcInputStats').textContent='0 символов';
        document.getElementById('tcOutputStats').textContent='—';
        document.getElementById('tcLog').classList.remove('active');
        tcStr='';showStatus('🗑️ Очищено!','success');
    }
    function tcDownload(){if(!tcStr){showStatus('⚠️ Сначала обработай!','error');return;} dlText(tcStr,(document.getElementById('tcFilename').value||'test_requests')+'.json','application/json');}
    function tcCopy(){if(!tcStr){showStatus('⚠️ Нечего копировать!','error');return;} cp(tcStr);}

    /* === TAB 6: COOKIE EXTRACTOR === */
    let ceResult='';
    document.getElementById('ceInput').addEventListener('input',function(){document.getElementById('ceInputStats').textContent=this.value.length+' символов';});

    function ceExtract(){
        const cookieString=document.getElementById('ceInput').value.trim();
        const cookieName=document.getElementById('ceCookieName').value.trim();
        const output=document.getElementById('ceOutput');

        if(!cookieString){showStatus('⚠️ Вставь строку cookie!','error');return;}
        if(!cookieName){showStatus('⚠️ Укажи имя cookie!','error');return;}

        const pairs=cookieString.split(';');
        let found='';

        for(let i=0;i<pairs.length;i++){
            let pair=pairs[i].trim();
            if(pair.startsWith(cookieName+'=')){
                found=pair.substring(cookieName.length+1);
                break;
            }
        }

        if(found){
            ceResult=found;
            output.textContent=found;
            output.style.color='#a6e3a1';
            document.getElementById('ceOutputStats').textContent=found.length+' символов';
            showStatus('✅ Cookie "'+cookieName+'" найдена!','success');
        }else{
            ceResult='';
            output.textContent='❌ Cookie "'+cookieName+'" не найдена в строке';
            output.style.color='#f38ba8';
            document.getElementById('ceOutputStats').textContent='—';
            showStatus('❌ Cookie "'+cookieName+'" не найдена','error');
        }
    }

    function ceClear(){
        document.getElementById('ceInput').value='';
        document.getElementById('ceOutput').innerHTML='<span style="color:#6c7086;font-style:italic;">Нажми "Извлечь", чтобы получить значение</span>';
        document.getElementById('ceOutputStats').textContent='—';
        document.getElementById('ceInputStats').textContent='0 символов';
        ceResult='';
        showStatus('🗑️ Очищено!','success');
    }

    function ceCopy(){
        if(!ceResult){showStatus('⚠️ Нечего копировать!','error');return;}
        cp(ceResult);
    }
    </script>
</body>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(htmlContent))
	})

	port := "8080"
	fmt.Println("==================================================")
	fmt.Println("✅ JSON Toolbox запущен! (6 инструментов)")
	fmt.Println("🌐 Открой браузер: http://localhost:" + port)
	fmt.Println("💡 Ctrl+C — остановить")
	fmt.Println("==================================================")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}