---
description: Create project plan using project-planner agent. No code writing - only plan file generation.
---

# /plan - Project Planning Mode

$ARGUMENTS

---

## 🔴 CRITICAL RULES

1. **NO CODE WRITING** - This command creates plan file only
2. **Use project-planner agent** - NOT Claude Code's native Plan subagent
3. **Socratic Gate** - Ask clarifying questions before planning
4. **Dynamic Naming** - Plan file named based on task
5. **Spec-Driven** - Plan matches `AGENTS.md` standards
6. **Data-First** - Plan must include Data Contracts (JSON Schema)

---

## Task

Use the `project-planner` agent with this context:

```
CONTEXT:
- User Request: $ARGUMENTS
- Mode: PLANNING ONLY (no code)
- Output: plans/{task-slug}/PLAN.md (Standard Location)

NAMING RULES:
1. Extract 2-3 key words from request
2. Lowercase, hyphen-separated
3. Max 30 characters
4. Example: "e-commerce cart" → plans/ecommerce-cart/PLAN.md

RULES:
1. Follow project-planner.md Phase -1 (Context Check)
2. Follow project-planner.md Phase 0 (Socratic Gate)
3. Create plans/{slug}/PLAN.md with task breakdown and Data Contracts
4. DO NOT write any code files
5. REPORT the exact file name created
```

---

## Expected Output

| Deliverable | Location |
|-------------|----------|
| Project Plan | `plans/{task-slug}/PLAN.md` |
| Task Breakdown | Inside plan file |
| Agent Assignments | Inside plan file |
| Verification Checklist | Phase X in plan file |

---

## After Planning

tell user:
```
[OK] Plan created: plans/{slug}/PLAN.md

Next steps:
- Review the plan
- Run `/create` to start implementation
- Or modify plan manually
```

---

## Naming Examples

| Request | Plan File |
|---------|-----------|
| `/plan e-commerce site with cart` | `plans/ecommerce-cart/PLAN.md` |
| `/plan mobile app for fitness` | `plans/fitness-app/PLAN.md` |
| `/plan add dark mode feature` | `plans/dark-mode/PLAN.md` |
| `/plan fix authentication bug` | `plans/auth-fix/PLAN.md` |
| `/plan SaaS dashboard` | `plans/saas-dashboard/PLAN.md` |

---

## Usage

```
/plan e-commerce site with cart
/plan mobile app for fitness tracking
/plan SaaS dashboard with analytics
```
