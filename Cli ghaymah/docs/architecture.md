```mermaid
classDiagram
    class GaymaaCLI {
        +Execute()
    }

    class Command {
        +Use: string
        +Short: string
        +Long: string
        +Run(cmd *Command, args []string)
    }

    class DeployCommand {
        -config: Config
        +Execute()
        -validateConfig()
        -uploadToGaymaa()
    }

    class StatusCommand {
        -appID: string
        +Execute()
        -getAppStatus()
        -displayStatus()
    }

    class LogsCommand {
        -appID: string
        -since: string
        +Execute()
        -fetchLogs()
        -displayLogs()
    }

    class GaymaaAPI {
        -baseURL: string
        -token: string
        +Deploy(config Config)
        +GetStatus(appID string)
        +GetLogs(appID string)
    }

    class Config {
        +AppName: string
        +DockerfilePath: string
        +EnvVars: map
    }

    GaymaaCLI --> Command
    Command <|-- DeployCommand
    Command <|-- StatusCommand
    Command <|-- LogsCommand
    DeployCommand --> GaymaaAPI
    StatusCommand --> GaymaaAPI
    LogsCommand --> GaymaaAPI
    DeployCommand --> Config
```

# Flow Diagram للعمليات الأساسية

```mermaid
sequenceDiagram
    participant User
    participant CLI
    participant GaymaaAPI
    participant Cloud

    %% Deploy Flow
    Note over User,Cloud: Deploy Flow
    User->>CLI: gaymaa deploy
    CLI->>CLI: Validate Config
    CLI->>GaymaaAPI: Upload Application
    GaymaaAPI->>Cloud: Deploy to Cloud
    Cloud-->>GaymaaAPI: Deployment Status
    GaymaaAPI-->>CLI: Status Response
    CLI-->>User: Show Deploy Status

    %% Status Flow
    Note over User,Cloud: Status Flow
    User->>CLI: gaymaa status
    CLI->>GaymaaAPI: Get App Status
    GaymaaAPI->>Cloud: Check Status
    Cloud-->>GaymaaAPI: Current Status
    GaymaaAPI-->>CLI: Status Info
    CLI-->>User: Display Status

    %% Logs Flow
    Note over User,Cloud: Logs Flow
    User->>CLI: gaymaa logs
    CLI->>GaymaaAPI: Request Logs
    GaymaaAPI->>Cloud: Fetch Logs
    Cloud-->>GaymaaAPI: App Logs
    GaymaaAPI-->>CLI: Log Data
    CLI-->>User: Display Logs
```

# المكونات الرئيسية

1. **GaymaaCLI**
   - المسؤول عن التعامل مع أوامر المستخدم
   - يدير الـ Commands المختلفة

2. **Commands**
   - `DeployCommand`: مسؤول عن رفع التطبيق
   - `StatusCommand`: مسؤول عن عرض حالة التطبيق
   - `LogsCommand`: مسؤول عن عرض السجلات

3. **GaymaaAPI**
   - يتعامل مع API غيمة
   - يدير عمليات الـ Deploy والـ Status والـ Logs

4. **Config**
   - يحتوي على إعدادات التطبيق
   - يدير متغيرات البيئة والإعدادات

# خطوات العمل

1. المستخدم يكتب أمر من خلال الـ CLI
2. الـ CLI يتحقق من صحة الأمر والمدخلات
3. يتم تنفيذ العملية المطلوبة من خلال GaymaaAPI
4. يتم عرض النتيجة للمستخدم
