# PermBAC 🛡️

Easy and flexible permission-based access control for Go applications.

PermBAC simplifies permission management in your Go projects, enabling detailed definition and efficient checking of permissions based on a structure that is configurable and adaptable to your needs.

## Installation

To install the PermBAC binary, use the following command:

```bash
go install https://github.com/eduardolat/permbac/cmd/permback
```

You can also download precompiled binaries from the [GitHub releases](https://github.com/eduardolat/permbac/releases).

## Getting Started

PermBAC operates through a clear and straightforward workflow, facilitating integration with your project:

1. **Initialization**: Generate a base configuration file with the command:

   ```bash
   permbac --init
   ```

2. **Permission Definition**: Write your permissions in the permission file. For example, in `permbac_perms.json`:

   ```json
   [
     {
       "name": "Example",
       "desc": "PermBAC permission example"
     },
     {
       "name": "...",
       "desc": "..."
     }
   ]
   ```

   You can see the detailed schema definition and validation rules [here.](internal/schema/perms.json)

3. **Permission Generation**: Use `permbac --generate` to create your permissions. You can specify a custom configuration file with:

   ```bash
   permbac --generate --config ./permbac.json
   ```

### Configuration File

The configuration file allows customization of the permission generation process. An example `permbac.json` could be:

```json
{
  "perms": ["./permbac_perms.json", "./perms/*.json"],
  "outdir": "./permbac",
  "package": "permbac"
}
```

This file defines the permission files to include (`perms`), the output directory for the generated files (`outdir`), and the package name (`package`).

You can see the detailed schema definition and validation rules [here.](internal/schema/config.json)

## Contributions

Contributions are welcome! Feel free to open an issue or pull request on [our repository](https://github.com/eduardolat/permbac).

## License

PermBAC is licensed under the [MIT License](https://opensource.org/licenses/MIT), allowing its broad use in both commercial and private projects.
