# Git Workflow & Commit Guidelines

GitVerse enforces strict Git hygiene, structured branch conventions, and [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) to keep the commit history clean, traceable, and automated.

---

## 1. Branch Strategy

* `main`: Protected production branch. All merged code must pass CI quality gates.
* Feature/Fix branches: Created off `main` and merged via Pull Requests.

### Branch Naming Conventions
* `feat/<short-description>`: New features or capabilities (e.g. `feat/galaxy-brain-counter`)
* `fix/<short-description>`: Bug fixes (e.g. `fix/profile-not-found-handling`)
* `docs/<short-description>`: Documentation changes (e.g. `docs/update-readme`)
* `perf/<short-description>`: Performance optimizations (e.g. `perf/parallel-badge-fetching`)
* `refactor/<short-description>`: Code restructuring without functional change

---

## 2. Conventional Commits Specification

Commit messages must follow the format:

```text
<type>(<optional scope>): <subject>

[optional body]

[optional footer(s)]
```

### Allowed Types
* `feat`: A new feature or capability.
* `fix`: A bug fix.
* `docs`: Documentation updates only.
* `style`: Styling, formatting, whitespace (no functional code changes).
* `refactor`: Code restructuring without bug fixes or new features.
* `perf`: Performance improvements.
* `test`: Adding or updating test suites.
* `ci`: Changes to CI/CD pipelines and workflows.
* `chore`: Maintenance tasks, dependency bumps.

### Examples
```bash
feat(calibration): support real-time Starstruck repository resolution
fix(parser): handle missing achievement timestamps on legacy accounts
perf(network): enable http keep-alive connection pooling
docs(readme): add dual-engine architecture diagram
```

---

## 3. Pull Request Guidelines

1. Link relevant issues in the PR description (`Closes #42`).
2. Keep PRs focused: avoid bundling unrelated changes into a single PR.
3. Ensure all tests pass locally before submitting:
   ```bash
   node --test internal/webui/jstest/*.test.mjs
   ```

