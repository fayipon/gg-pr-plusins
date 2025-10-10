```mermaid
graph TD

subgraph G["🌍 API Gateway (Nginx / Kong)"]
  G1[/api/promotion/]
  G2[/api/decision/]
  G3[/api/billing/]
end

subgraph A["🎯 Promotion Service Cluster"]
  A1[promotion-api<br>活動管理 + 可見控制]
  A2[promotion-engine<br>JSONLogic + Formula 運算]
  A3[promotion-worker<br>發獎 / 回調任務]
  A4[promotion-template<br>模板中心]
  A5[promotion-editor<br>可視化編輯器]
end

subgraph B["📊 Data Decision Service Cluster"]
  B1[metric-engine<br>KPI 統計引擎]
  B2[feature-engine<br>玩家特徵建模]
  B3[decision-engine<br>策略決策 + 分群模型]
  B4[feedback-engine<br>活動結果學習與調整]
end

subgraph C["💰 Billing & Budget Cluster"]
  C1[budget-service<br>租戶預算管理]
  C2[usage-service<br>API 用量統計]
  C3[billing-service<br>計費與警報]
end

subgraph D["📢 Marketing Automation Cluster"]
  D1[trigger-service<br>行為事件觸發器]
  D2[abtest-service<br>A/B 測試與流量控制]
  D3[locale-service<br>多語系模板中心]
  D4[notify-service<br>通知與Webhook中心]
end

subgraph E["📈 Analytics Cluster"]
  E1[report-service<br>活動報表與ROI分析]
  E2[datalake-service<br>數據倉庫 / ClickHouse]
end

subgraph F["⚙️ Shared Infra"]
  F1[(PostgreSQL / TiDB)]
  F2[(Redis Cache)]
  F3[(Kafka / NATS MQ)]
  F4[(Prometheus / Grafana)]
end

G1 --> A1
A1 --> A2 --> A3 --> D4
A2 --> B3
A3 --> C1
B1 --> B3 --> B4 --> A2
B4 --> D1
D1 --> A1
D2 --> A1
D3 --> D4
E1 --> B4
A1 --> E1

A1 & A2 & A3 & B1 & B3 & C1 & D1 & D2 --> F1
F1 & F2 & F3 --> F4
```

| 類別                    | 微服務名稱                | 核心職責             | 技術要點                        |
| --------------------- | -------------------- | ---------------- | --------------------------- |
| **Promotion Service** | `promotion-api`      | 活動 CRUD、可見條件     | JWT + gRPC + JSONLogic      |
|                       | `promotion-engine`   | 參與 / 符合 / 奬勵公式運算 | Govaluate + Lua             |
|                       | `promotion-worker`   | 發獎與回調任務          | MQ 消費 + 錢包 API              |
|                       | `promotion-template` | 模板中心             | JSON Schema + 快速生成          |
|                       | `promotion-editor`   | 視覺化邏輯編輯器         | Blockly / Vue Builder       |
| **Decision Service**  | `metric-engine`      | 指標統計：ROI / 留存    | ClickHouse / SQL 聚合         |
|                       | `feature-engine`     | 玩家特徵嵌入建模         | Python + ML Pipeline        |
|                       | `decision-engine`    | 策略決策與分群          | XGBoost / JSONLogic         |
|                       | `feedback-engine`    | 活動結果回饋訓練         | Auto-Tuning / Reinforcement |
| **Billing Service**   | `budget-service`     | 預算上限 + 鎖單        | Redis 鎖定 / 監控警報             |
|                       | `usage-service`      | API 使用統計         | gRPC Middleware             |
|                       | `billing-service`    | 成本計算與通知          | CRON + Telegram Alert       |
| **Marketing Service** | `trigger-service`    | 行為事件觸發           | Kafka / Event Bus           |
|                       | `abtest-service`     | A/B 測試與流量控制      | 分流演算法 / Redis Hash          |
|                       | `locale-service`     | 多語模板             | i18n JSON + 動態替換            |
|                       | `notify-service`     | 通知中心             | OneSignal / SMTP / Webhook  |
| **Analytics Service** | `report-service`     | 成效報表與ROI分析       | Grafana / CSV Export        |
|                       | `datalake-service`   | 數據倉庫整合           | TiDB + ClickHouse ETL       |


資料流與事件流設計
📥 Data Flow

promotion-api 收到活動事件

promotion-engine 進行條件運算

worker 將結果寫入 MQ

decision-engine 同步更新 KPI 模型

billing-service 計算成本與預算消耗

report-service 彙總指標 → 供 dashboard 顯示

topic: promotion.event.join
topic: promotion.event.reward
topic: decision.update.model
topic: billing.budget.alert
topic: trigger.user.deposit
